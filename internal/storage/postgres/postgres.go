package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/logger"
)

// PostgresStorage реализует хранилище URL в PostgreSQL
type PostgresStorage struct {
	db *sql.DB // Подключение к базе данных
}

// NewPostgresStorage создает новый экземпляр хранилища PostgreSQL
// dsn - строка подключения к базе данных
// Выполняет подключение к базе данных и применяет миграции
func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := connect(dsn)
	if err != nil {
		return nil, err
	}

	if err = migrateDB(dsn); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

// Save сохраняет пару сокращенный URL - оригинальный URL в базе данных
// ctx - контекст запроса
// shortID - сокращенный URL
// originalURL - оригинальный URL
// userID - идентификатор пользователя
// Возвращает ErrURLAlreadyExists, если URL уже существует
func (ps *PostgresStorage) Save(ctx context.Context, shortID, originalURL, userID string) error {
	_, err := ps.db.ExecContext(ctx,
		`INSERT INTO urls (short_url, original_url, user_id) VALUES ($1, $2, $3)`,
		shortID, originalURL, userID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return ErrURLAlreadyExists
		}
		return err
	}

	return nil
}

// SaveBatch сохраняет пакет URL в базе данных
// ctx - контекст запроса
// urls - карта сокращенных URL к оригинальным URL
// userID - идентификатор пользователя
// При конфликте по short_url операция пропускается
func (ps *PostgresStorage) SaveBatch(ctx context.Context, urls map[string]string, userID string) error {
	if len(urls) == 0 {
		return nil
	}

	var (
		builder strings.Builder
		params  []interface{}
		i       = 1
	)

	builder.WriteString("INSERT INTO urls (short_url, original_url, user_id) VALUES ")

	for short, orig := range urls {
		builder.WriteString(fmt.Sprintf("($%d, $%d, $%d),", i, i+1, i+2))
		params = append(params, short, orig, userID)
		i += 3
	}

	query := strings.TrimSuffix(builder.String(), ",") + " ON CONFLICT (short_url) DO NOTHING"

	_, err := ps.db.ExecContext(ctx, query, params...)
	if err != nil {
		logger.Log.Error("Failed batch insert", zap.Error(err))
		return fmt.Errorf("batch insert failed: %w", err)
	}

	return nil
}

// Get возвращает оригинальный URL по сокращенному
// ctx - контекст запроса
// shortID - сокращенный URL
// Возвращает:
// - оригинальный URL
// - флаг удаления
// - флаг наличия URL в хранилище
func (ps *PostgresStorage) Get(ctx context.Context, shortID string) (string, bool, bool) {
	var originalURL string
	var isDeleted bool

	err := ps.db.QueryRowContext(
		ctx,
		`SELECT original_url, is_deleted FROM urls WHERE short_url = $1 ORDER BY id DESC LIMIT 1`,
		shortID,
	).Scan(&originalURL, &isDeleted)

	if err == sql.ErrNoRows {
		return "", false, false
	} else if err != nil {
		logger.Log.Error("Error of get url from db", zap.Error(err))
		return "", false, false
	}

	return originalURL, isDeleted, true
}

// GetUserURLs возвращает список URL пользователя
// ctx - контекст запроса
// uuid - идентификатор пользователя
// Возвращает список пар сокращенный URL - оригинальный URL
func (ps *PostgresStorage) GetUserURLs(ctx context.Context, uuid string) ([]dto.UserURLItem, error) {
	rows, err := ps.db.QueryContext(ctx,
		`SELECT short_url, original_url FROM urls WHERE user_id = $1`, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.UserURLItem
	for rows.Next() {
		var item dto.UserURLItem
		if err := rows.Scan(&item.ShortURL, &item.OriginalURL); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteUserURLs помечает URL пользователя как удаленные
// ctx - контекст запроса
// userID - идентификатор пользователя
// uuids - список сокращенных URL для удаления
func (ps *PostgresStorage) DeleteUserURLs(ctx context.Context, userID string, uuids []string) error {
	if len(uuids) == 0 {
		return nil
	}

	var (
		params  = make([]interface{}, 0, len(uuids)+1)
		builder strings.Builder
	)

	params = append(params, userID)
	builder.WriteString("UPDATE urls SET is_deleted = true WHERE user_id = $1 AND short_url IN (")

	for i := range uuids {
		builder.WriteString(fmt.Sprintf("$%d,", i+2))
		params = append(params, uuids[i])
	}

	query := strings.TrimSuffix(builder.String(), ",") + ")"
	_, err := ps.db.ExecContext(ctx, query, params...)
	return err
}

// FindShortByOriginalURL ищет сокращенный URL по оригинальному
// ctx - контекст запроса
// originalURL - оригинальный URL
// Возвращает сокращенный URL или ошибку, если URL не найден
func (ps *PostgresStorage) FindShortByOriginalURL(ctx context.Context, originalURL string) (string, error) {
	var short string
	err := ps.db.QueryRowContext(ctx,
		`SELECT short_url FROM urls WHERE original_url = $1 LIMIT 1`,
		originalURL).Scan(&short)

	if err != nil {
		return "", err
	}
	return short, nil
}

// GetStats возвращает статистику по хранилищу URL
// ctx - контекст запроса
// Возвращает:
// - количество активных (не удаленных) URL
// - количество уникальных пользователей
// - ошибку, если произошла при выполнении запросов
func (ps *PostgresStorage) GetStats(ctx context.Context) (int, int, error) {
	var urlCount, userCount int
	err := ps.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM urls WHERE is_deleted = false`).Scan(&urlCount)
	if err != nil {
		return 0, 0, err
	}

	err = ps.db.QueryRowContext(ctx, `SELECT COUNT(DISTINCT user_id) FROM urls WHERE user_id IS NOT NULL`).Scan(&userCount)
	return urlCount, userCount, err
}

// Ping проверяет соединение с базой данных
// ctx - контекст для выполнения запроса
func (ps *PostgresStorage) Ping(ctx context.Context) error {
	return ps.db.PingContext(ctx)
}

// Close закрывает соединение с базой данных
func (ps *PostgresStorage) Close() error {
	return ps.db.Close()
}

// connect устанавливает соединение с базой данных PostgreSQL
// dsn - строка подключения к базе данных
func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres db: %w", err)
	}

	logger.Log.Info("Successfully connected to postgres database")

	return db, nil
}

// migrateDB применяет миграции к базе данных
// dsn - строка подключения к базе данных
func migrateDB(dsn string) error {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return fmt.Errorf("failed to create db migration: %w", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("failed to apply db migration: %w", err)
	}

	logger.Log.Info("Successfully applied db migration")

	return nil
}

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

type PostgresStorage struct {
	db *sql.DB
}

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

func (ps *PostgresStorage) Save(shortID, originalURL, userID string) error {
	_, err := ps.db.ExecContext(context.Background(),
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

func (ps *PostgresStorage) SaveBatch(urls map[string]string, userID string) error {
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

	_, err := ps.db.ExecContext(context.Background(), query, params...)
	if err != nil {
		logger.Log.Error("Failed batch insert", zap.Error(err))
		return fmt.Errorf("batch insert failed: %w", err)
	}

	return nil
}

func (ps *PostgresStorage) Get(shortID string) (string, bool) {
	var originalURL string
	err := ps.db.QueryRowContext(
		context.Background(),
		`SELECT original_url FROM urls WHERE short_url = $1 ORDER BY id DESC LIMIT 1`,
		shortID,
	).Scan(&originalURL)

	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		logger.Log.Error("Error of open file", zap.Error(err))
		return "", false
	}

	return originalURL, true
}

func (ps *PostgresStorage) GetUserURLs(uuid string) ([]dto.UserURLItem, error) {
	rows, err := ps.db.QueryContext(context.Background(),
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

func (ps *PostgresStorage) FindShortByOriginalURL(originalURL string) (string, error) {
	var short string
	err := ps.db.QueryRowContext(context.Background(),
		`SELECT short_url FROM urls WHERE original_url = $1 LIMIT 1`,
		originalURL).Scan(&short)

	if err != nil {
		return "", err
	}
	return short, nil
}

func (ps *PostgresStorage) Ping(ctx context.Context) error {
	return ps.db.PingContext(ctx)
}

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

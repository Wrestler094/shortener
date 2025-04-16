package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

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

func (ps *PostgresStorage) Save(shortID string, originalURL string) {
	_, _ = ps.db.ExecContext(
		context.Background(),
		`INSERT INTO urls (short_id, original_url) VALUES ($1, $2) ON CONFLICT (short_id) DO NOTHING`,
		shortID, originalURL,
	)
}

func (ps *PostgresStorage) Get(shortID string) (string, bool) {
	var originalURL string
	err := ps.db.QueryRowContext(
		context.Background(),
		`SELECT original_url FROM urls WHERE short_id = $1 ORDER BY id DESC LIMIT 1`,
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

	log.Println("Connected to PostgreSQL using pgx.")

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

	log.Println("Migrations applied successfully.")

	return nil
}

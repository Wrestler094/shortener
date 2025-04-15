package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/golang-migrate/migrate/v4"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitPostgres(dsn string) error {
	if configs.FlagDatabaseDSN == "" {
		return nil
	}

	var err error

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres db: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping postgres db: %w", err)
	}

	log.Println("Connected to PostgreSQL using pgx.")

	if err = Migrate(dsn); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func Migrate(dsn string) error {
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

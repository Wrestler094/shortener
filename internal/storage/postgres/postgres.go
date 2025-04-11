package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitPostgres(dsn string) error {
	var err error

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres db: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping postgres db: %w", err)
	}

	log.Println("Connected to PostgreSQL using pgx.")
	return nil
}

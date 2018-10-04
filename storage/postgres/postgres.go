package postgres

import (
	"database/sql"
	"fmt"
)

// Postgres implements Storage interface
type Postgres struct {
	db *sql.DB
}

// Keep implements Storage.Keep method
func (psql *Postgres) Keep(short, longURL string) error {
	query := "INSERT INTO shortener (short, long) VALUES ($1, $2);"
	_, err := psql.db.Exec(query, short, longURL)
	if err != nil {
		return fmt.Errorf("execute query %s: %v", query, err)
	}
	return nil
}

// Extract implements Storage.Extract method
func (psql *Postgres) Extract(short string) (string, error) {
	longURL := ""
	query := "SELECT long FROM shortener WHERE short = $1;"
	row := psql.db.QueryRow(query, short)
	if err := row.Scan(&longURL); err != nil {
		return "", fmt.Errorf("scan %s stmt: %v", query, err)
	}
	return longURL, nil
}

// New returns the new Postgres object
func New(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

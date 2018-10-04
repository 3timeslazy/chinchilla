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
	rows, err := psql.db.Query(query, short)
	if err != nil {
		return "", fmt.Errorf("query %s: %v", query, err)
	}

	// do not use method db.QueryRow above because
	// github.com/lib/pq driver doesn't return an error
	// when there are no sql.Rows
	// so we do it for the driver
	if !rows.Next() {
		return "", sql.ErrNoRows
	}

	err = rows.Scan(&longURL)
	if err != nil {
		return "", fmt.Errorf("scan long url: %v", err)
	}

	return longURL, nil
}

// New returns the new Postgres object
func New(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

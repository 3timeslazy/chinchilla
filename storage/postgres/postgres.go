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

// GetLongByShort returns long url by short
func (psql *Postgres) GetLongByShort(short string) (string, error) {
	longURL := ""
	query := "SELECT long FROM shortener WHERE short = $1;"
	rows, err := psql.db.Query(query, short)
	if err != nil {
		return "", fmt.Errorf("query %s: %v", query, err)
	}

	if !rows.Next() {
		return "", nil
	}

	err = rows.Scan(&longURL)
	if err != nil {
		return "", fmt.Errorf("scan long url: %v", err)
	}

	return longURL, nil
}

// GetShortByLong returns short url by long
func (psql *Postgres) GetShortByLong(longURL string) (string, error) {
	query := "SELECT short FROM shortener WHERE long = $1;"
	rows, err := psql.db.Query(query, longURL)
	if err != nil {
		return "", fmt.Errorf("query %s: %v", query, err)
	}

	if !rows.Next() {
		return "", nil
	}

	var short string
	if err = rows.Scan(&short); err != nil {
		return "", fmt.Errorf("scan sql.Row: %v", err)
	}

	return short, nil
}

// New returns the new Postgres object
func New(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

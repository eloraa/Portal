package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(url string) (*Database, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) CreateUser() (string, error) {
	var userID string
	err := d.db.QueryRow(`
		INSERT INTO users DEFAULT VALUES
		RETURNING id
	`).Scan(&userID)
	return userID, err
}

func (d *Database) InitializeTables() error {
	_, err := d.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid()
		)
	`)
	return err
}

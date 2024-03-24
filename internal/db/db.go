package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents the database connection.
type DB struct {
	Conn *pgxpool.Pool
}

// NewDB creates a new DB instance and connects to the database.
func NewDB() (*DB, error) {
	conn, err := pgxpool.New(context.Background(), "postgresql://user:pass@localhost:5432/db")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return &DB{Conn: conn}, nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	db.Conn.Close()
	return nil
}

// ExecuteQuery executes the given SQL query on the database.
func (db *DB) ExecuteQuery(query string) error {
	_, err := db.Conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

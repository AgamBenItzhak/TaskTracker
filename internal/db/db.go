package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB initializes the database connection pool
func InitDB(ctx context.Context, dsn string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return dbpool
}

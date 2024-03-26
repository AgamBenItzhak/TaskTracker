package db

import (
	"context"
	"log"

	"github.com/AgamBenItzhak/TaskTracker/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB provides functions for the database
type DB struct {
	Pool *pgxpool.Pool
}

func NewDB() (*DB, error) {
	dbHost := config.GetString("db.host")
	dbPort := config.GetString("db.port")
	dbUser := config.GetString("db.username")
	dbPassword := config.GetString("db.password")
	dbName := config.GetString("db.database")

	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DB{Pool: pool}, nil
}

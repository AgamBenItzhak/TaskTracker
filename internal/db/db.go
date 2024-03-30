package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AgamBenItzhak/TaskTracker/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxIface interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Config() *pgxpool.Config
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Reset()
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Stat() *pgxpool.Stat
}

// DB provides functions for the database
type DB struct {
	Pool *pgxpool.Pool
}

func NewDB() (DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.ServerConfig.DB.User,
		config.ServerConfig.DB.Password,
		config.ServerConfig.DB.Host,
		config.ServerConfig.DB.Port,
		config.ServerConfig.DB.DBName)

	if config.ServerConfig.DB.SSLMode != "" {
		dsn += fmt.Sprintf("?sslmode=%s", config.ServerConfig.DB.SSLMode)
	}

	if err := initializeDB(dsn); err != nil {
		return DB{}, err
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return DB{}, err
	}

	return DB{Pool: pool}, nil
}

func initializeDB(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	// Check if the database is up to date and not dirty
	currentVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return err
	}

	if dirty {
		return fmt.Errorf("database is dirty, manual intervention required. current version: %d", currentVersion)
	}

	// if currentVersion == 0 || err == migrate.ErrNilVersion {
	// 	return fmt.Errorf("database is not initialized, applying migrations")
	// }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

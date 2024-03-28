package db

import (
	"context"
	"log"

	"github.com/AgamBenItzhak/TaskTracker/config"

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
	dbHost := config.GetString("db.host")
	dbPort := config.GetString("db.port")
	dbUser := config.GetString("db.username")
	dbPassword := config.GetString("db.password")
	dbName := config.GetString("db.database")

	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	return DB{Pool: pool}, nil
}

package repository

import (
	"context"
	"database/sql"
)

type SqlMethods interface {
	InitDb(or, wr string) error
	Close() error

	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	BeginTransaction() (*sql.Tx, error)
	CommitTransaction(tx *sql.Tx) error
	RollbackTransaction(tx *sql.Tx) error
}

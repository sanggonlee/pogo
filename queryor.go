package pogo

import (
	"context"
	"database/sql"
)

// Queryor is an interface for general Postgres connection pool.
// sql.DB, sql.Stmt, sql.Tx, and other extensions to the Go sql package
// provide abstractions that implement this.
type Queryor interface {
	// Query is used to query the database.
	Query(query string, args ...interface{}) (*sql.Rows, error)

	// QueryContext is used to query the database with a context.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

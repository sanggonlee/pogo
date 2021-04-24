package pogo

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/sanggonlee/pogo/internal/query"
	"github.com/sanggonlee/pogo/internal/version"
)

// Query prepares a query running instance. Note that it does not run a
// query by itself.
func Query(queryor Queryor) QueryRunner {
	safeguardPostgresVersion()
	return QueryRunner{
		queryor: queryor,
	}
}

// QueryContext is like Query, except it also takes a context.
func QueryContext(ctx context.Context, queryor Queryor) QueryRunner {
	safeguardPostgresVersion()
	return QueryRunner{
		queryor: queryor,
		ctx:     ctx,
	}
}

// QueryRunner is an encapsulation for running a single query.
type QueryRunner struct {
	queryor Queryor
	ctx     context.Context
}

// For is used to run a query on arbitrary relations. It returns sql.Rows, which
// you can scan to the structure you see fit.
func (qr QueryRunner) For(queryable query.Queryable) (*sql.Rows, error) {
	q, err := queryable.ToQuery()
	if err != nil {
		return nil, errors.Wrap(err, "converting queryable to query string")
	}

	var rows *sql.Rows
	if qr.ctx == nil {
		rows, err = qr.queryor.Query(q)
	} else {
		rows, err = qr.queryor.QueryContext(qr.ctx, q)
	}
	if err != nil {
		return nil, errors.Wrap(err, "querying rows")
	}

	return rows, nil
}

func safeguardPostgresVersion() {
	// Ensure Postgres version is locked down before running any queries.
	if !version.IsSet() {
		_ = SetPostgresVersion(defaultPostgresVersion)
	}
}

package pogo_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/sanggonlee/pogo"
)

func ExampleQuery() {
	var db pogo.Queryor
	rows, _ := pogo.Query(db).For(
		pogo.StatDatabaseView.With(
			pogo.StatActivityView.With(
				pogo.BlockingPIDs,
			),
			pogo.LocksView,
		),
	)
	defer rows.Close()
	for rows.Next() {
		// ...
	}
}

func TestQueryRunner_For(t *testing.T) {
	cases := []struct {
		description string
		queryor     pogo.Queryor
		ctx         context.Context
	}{
		{
			description: "Query should run if context is nil",
			queryor:     &mockQueryor{queryContextError: errors.New("this should not have been run")},
		},
		{
			description: "QueryContext should run if context is not nil",
			queryor:     &mockQueryor{queryError: errors.New("this should not have been run")},
			ctx:         context.Background(),
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			qr := pogo.Query(c.queryor)
			if c.ctx != nil {
				qr = pogo.QueryContext(c.ctx, c.queryor)
			}

			_, err := qr.For(pogo.LocksView)
			if err != nil {
				t.Errorf("Expected nil error but got %v", err)
			}
		})
	}
}

type mockQueryor struct {
	queryError        error
	queryContextError error
}

func (mq mockQueryor) Query(q string, args ...interface{}) (*sql.Rows, error) {
	if mq.queryError != nil {
		return nil, mq.queryError
	}
	return &sql.Rows{}, nil
}

func (mq mockQueryor) QueryContext(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	if mq.queryContextError != nil {
		return nil, mq.queryContextError
	}
	return &sql.Rows{}, nil
}

package pogo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3"
)

// TupleArgs specifies the arguments to feed when querying for a tuple.
type TupleArgs struct {
	// Name of the relation. Corresponds to "relname" column in pg_stat_user_tables view.
	RelName string

	// Page of the tuple to query for.
	Page int64

	// Tuple index of the tuple to query for.
	Tuple int64
}

// Tuple runs a query on a specific relation matching the ctid given by a pair (page,tuple)
// and returns the resulting row as a JSON string, with its properties mapping to the columns.
// Note that tuples retrieved this way are no longer consistent after VACUUM FULL.
func (qr QueryRunner) Tuple(ctx context.Context, args TupleArgs) (null.String, error) {
	var rowJSON null.String
	q := fmt.Sprintf(`
		SELECT row_to_json(%[1]s)
		FROM %[1]s
		WHERE ctid = '(%d,%d)'
	`, args.RelName, args.Page, args.Tuple)
	rows, err := qr.queryor.QueryContext(ctx, q)
	if err != nil {
		return rowJSON, errors.Wrap(err, "querying tuple")
	}

	for rows.Next() {
		if rowJSON.Valid {
			// Only one row should've been fetched
			return rowJSON, errors.Wrap(err, "page and tuple state got out of sync")
		}

		if err := rows.Scan(&rowJSON); err != nil {
			return rowJSON, errors.Wrap(err, "scanning tuple")
		}
	}

	return rowJSON, nil
}

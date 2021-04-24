package pogo

import (
	"github.com/pkg/errors"
	"github.com/sanggonlee/pogo/internal/query"
	"github.com/sanggonlee/pogo/postgres9"
)

// StatActivity9 is a convenience method for running a query on pg_stat_activity view.
// It is meant to be used for Postgres v9.6.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatActivity9(where string, joins ...query.Queryable) ([]postgres9.StatActivityJoined, error) {
	if v := GetPostgresVersion(); v != Postgres9 {
		return nil, getVersionMismatchError(Postgres9, v)
	}

	queryable := StatActivityView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_activity")
	}
	defer rows.Close()

	ss := make([]postgres9.StatActivityJoined, 0)
	for rows.Next() {
		var s postgres9.StatActivityJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_activity row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// StatReplication9 is a convenience method for running a query on pg_stat_replication view.
// It is meant to be used for Postgres v9.6.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatReplication9(where string, joins ...query.Queryable) ([]postgres9.StatReplicationJoined, error) {
	if v := GetPostgresVersion(); v != Postgres9 {
		return nil, getVersionMismatchError(Postgres9, v)
	}

	queryable := StatReplicationView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_replication")
	}
	defer rows.Close()

	ss := make([]postgres9.StatReplicationJoined, 0)
	for rows.Next() {
		var s postgres9.StatReplicationJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_replication row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// StatTable9 is a convenience method for running a query on pg_stat_user_tables view.
// It is meant to be used for Postgres v9.6.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatTable9(where string, joins ...query.Queryable) ([]postgres9.StatTableJoined, error) {
	if v := GetPostgresVersion(); v != Postgres9 {
		return nil, getVersionMismatchError(Postgres9, v)
	}

	queryable := StatUserTablesView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_user_tables")
	}
	defer rows.Close()

	ss := make([]postgres9.StatTableJoined, 0)
	for rows.Next() {
		var s postgres9.StatTableJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_user_tables row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// Locks9 is a convenience method for running a query on pg_locks view.
// It is meant to be used for Postgres v9.6.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) Locks9(where string, joins ...query.Queryable) ([]postgres9.LockJoined, error) {
	if v := GetPostgresVersion(); v != Postgres9 {
		return nil, getVersionMismatchError(Postgres9, v)
	}

	queryable := LocksView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_locks")
	}
	defer rows.Close()

	ls := make([]postgres9.LockJoined, 0)
	for rows.Next() {
		var l postgres9.LockJoined
		dest := l.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_locks row")
		}

		ls = append(ls, l)
	}

	return ls, nil
}

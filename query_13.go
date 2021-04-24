package pogo

import (
	"github.com/pkg/errors"
	"github.com/sanggonlee/pogo/internal/query"
	"github.com/sanggonlee/pogo/postgres13"
)

// StatActivity13 is a convenience method for running a query on pg_stat_activity view.
// It is meant to be used for Postgres v13.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatActivity13(where string, joins ...query.Queryable) ([]postgres13.StatActivityJoined, error) {
	if v := GetPostgresVersion(); v != Postgres13 {
		return nil, getVersionMismatchError(Postgres13, v)
	}

	queryable := StatActivityView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_activity")
	}
	defer rows.Close()

	ss := make([]postgres13.StatActivityJoined, 0)
	for rows.Next() {
		var s postgres13.StatActivityJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_activity row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// StatReplication13 is a convenience method for running a query on pg_stat_replication view.
// It is meant to be used for Postgres v13.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatReplication13(where string, joins ...query.Queryable) ([]postgres13.StatReplicationJoined, error) {
	if v := GetPostgresVersion(); v != Postgres13 {
		return nil, getVersionMismatchError(Postgres13, v)
	}

	queryable := StatReplicationView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_replication")
	}
	defer rows.Close()

	ss := make([]postgres13.StatReplicationJoined, 0)
	for rows.Next() {
		var s postgres13.StatReplicationJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_replication row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// StatTable13 is a convenience method for running a query on pg_stat_user_tables view.
// It is meant to be used for Postgres v13.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) StatTable13(where string, joins ...query.Queryable) ([]postgres13.StatTableJoined, error) {
	if v := GetPostgresVersion(); v != Postgres13 {
		return nil, getVersionMismatchError(Postgres13, v)
	}

	queryable := StatUserTablesView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_stat_user_tables")
	}
	defer rows.Close()

	ss := make([]postgres13.StatTableJoined, 0)
	for rows.Next() {
		var s postgres13.StatTableJoined
		dest := s.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_stat_user_tables row")
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// Locks13 is a convenience method for running a query on pg_locks view.
// It is meant to be used for Postgres v13.
// If you want to select rows with certain conditions, pass a non-empty where argument,
// which will be injected as WHERE {where} in the query.
func (qr QueryRunner) Locks13(where string, joins ...query.Queryable) ([]postgres13.LockJoined, error) {
	if v := GetPostgresVersion(); v != Postgres13 {
		return nil, getVersionMismatchError(Postgres13, v)
	}

	queryable := LocksView.
		Where(where).
		With(joins...)

	rows, err := qr.For(queryable)
	if err != nil {
		return nil, errors.Wrap(err, "querying pg_locks")
	}
	defer rows.Close()

	ls := make([]postgres13.LockJoined, 0)
	for rows.Next() {
		var l postgres13.LockJoined
		dest := l.ScanDestinations(queryable.Joins)

		if err := rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "scanning pg_locks row")
		}

		ls = append(ls, l)
	}

	return ls, nil
}

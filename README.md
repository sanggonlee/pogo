pogo
==================

[![Go Reference](https://pkg.go.dev/badge/github.com/sanggonlee/pogo.svg)](https://pkg.go.dev/github.com/sanggonlee/pogo)
[![Go Report Card](https://goreportcard.com/badge/github.com/sanggonlee/pogo)](https://goreportcard.com/report/github.com/sanggonlee/pogo)


pogo is a lightweight Go PostgreSQL internal state query engine.

It focuses on the data that are highly dynamic in nature, and provides some convenience on recursivley joining system catalogue relations and function calls on the fly.

Was mainly created to power [plum](https://github.com/sanggonlee/plum) project (along with [monochron](https://github.com/sanggonlee/monochron)).

pogo supports querying on the following relations:
```
pg_locks
pg_stat_activity
pg_stat_replication
pg_stat_ssl
pg_stat_gssapi (for Postgres 13)
pg_stat_wal_receiver
pg_stat_subscription (for Postgres 13)
pg_stat_database
pg_stat_database_conflicts
pg_stat_user_tables
pg_stat_user_indexes
pg_statio_user_indexes
pg_statio_user_sequences
pg_statio_user_tables
pg_stat_user_tables
pg_stat_user_functions
pg_stat_archiver
pg_stat_bgwriter
pg_stat_slru (for Postgres 13)
```

Currently only supports PostgreSQL 9.6 and 13.

## Usage
```
rows, err := pogo.Query(sql.DB).For(
 	pogo.StatDatabaseView.With(
 		pogo.StatActivityView.With(
			pogo.BlockingPIDs
  		),
  		pogo.LocksView,
  	),
)
```

will return `sql.Rows` which you can scan to the appropriate objects. In this case it will be all rows in `pg_stat_database`, left joined with `pg_locks` and `pg_stat_activity` on `datid`, as well as a list of `pg_blocking_pids(pg_stat_activity.pid)` under each row of `pg_stat_activity`.
You can find the struct definitions under `postgres9` and `postgres13` subpackages. Please refer to the godoc.

## Documentation

[godoc](https://pkg.go.dev/github.com/sanggonlee/pogo)
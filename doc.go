//

// Pogo is a lightweight Go library for querying a subset of PostgreSQL's internal states.
// It focuses on the data that are highly dynamic in nature, and provides some convenience
// on recursivley joining some system catalogue relations and function calls on the fly.
// Currently it supports PostgreSQL version 9.6 and 13 only.
//
// In particular, it supports querying on the following relations:
//   pg_locks
//   pg_stat_activity
//   pg_stat_replication
//   pg_stat_ssl
//   pg_stat_gssapi
//   pg_stat_wal_receiver
//   pg_stat_subscription
//   pg_stat_database
//   pg_stat_database_conflicts
//   pg_stat_user_tables
//   pg_stat_user_indexes
//   pg_statio_user_indexes
//   pg_statio_user_sequences
//   pg_statio_user_tables
//   pg_stat_user_tables
//   pg_stat_user_functions
//   pg_stat_archiver
//   pg_stat_bgwriter
//   pg_stat_slru
//
// Among these relations, the following four relations are supported as "primary" query targets
// (i.e. pogo provides functions that query them directly and give you the
// resulting rows as defined structs)
//   pg_locks
//   pg_stat_activity
//   pg_stat_replication
//   pg_stat_user_tables
//
// Examples:
//
// Querying relations other than the four primary supported targets
// (querying pg_stat_database view joining with pg_stat_activity view and
// pg_locks view, and pg_stat_activity view in turn includes pg_blocking_pids(pid)):
//  rows, err := pogo.Query(sql.DB).For(
// 	pogo.StatDatabaseView.With(
// 		pogo.StatActivityView.With(
//			pogo.BlockingPIDs
//  		),
//  		pogo.LocksView,
//  	),
//  )
//
// Querying pg_stat_activity for Postgres 13, joining with pg_locks:
//  statActivities, err := pogo.Query(sql.DB).StatActivity13(
// 	pogo.LocksView,
//  )
//
// The currently supported joins are as follows:
//
//   pg_stat_activity
//    - pg_locks (on pid)
//    - pg_locks (on backend_xid)
//    - pg_stat_ssl (on pid)
//    - pg_stat_gssapi (on pid)
//    - pg_stat_wal_receiver (on pid)
//    - pg_stat_subscription (on pid)
//    - pg_stat_database (on pid)
//    - pg_stat_database_conflicts (on pid)
//    - pg_blocking_pids(pg_stat_activity.pid)
//   pg_stat_replication
//    - pg_locks (on pid)
//    - pg_stat_ssl (on pid)
//    - pg_stat_gssapi (on pid)
//    - pg_stat_wal_receiver (on pid)
//   pg_stat_user_tables
//    - pg_locks (on relid)
//    - pg_stat_user_indexes (on relid)
//    - pg_stat_subscription (on relid)
//    - pg_statio_user_indexes (on relid)
//    - pg_statio_user_sequences (on relid)
//    - pg_statio_user_tables (on relid)
//   pg_locks
//    - pg_stat_activity (on pid)
//    - pg_stat_database (on database)
//    - pg_stat_user_tables (on relation)
//    - pg_stat_user_indexes (on relation)
//    - pg_statio_user_tables (on relation)
//    - pg_statio_user_indexes (on relation)
//    - pg_statio_user_sequences (on relation)
//   pg_stat_ssl
//    - pg_locks (on pid)
//    - pg_stat_activity (on pid)
//   pg_stat_gssapi
//    - pg_locks (on pid)
//    - pg_stat_activity (on pid)
//   pg_stat_wal_receiver
//    - pg_locks (on pid)
//    - pg_stat_activity (on pid)
//   pg_stat_subscription
//    - pg_locks (on pid)
//    - pg_stat_activity (on pid)
//   pg_stat_database
//    - pg_database_conflicts (on datid)
//    - pg_locks (on datid)
//    - pg_stat_activity (on datid)
//   pg_stat_database_conflicts
//   pg_stat_user_indexes
//    - pg_stat_user_tables (on relid)
//    - pg_statio_user_tables (on relid)
//    - pg_locks (on indexrelid)
//    - pg_statio_user_indexes (on indexrelid)
//   pg_statio_user_indexes
//   pg_statio_user_sequences
//   pg_statio_user_tables
//   pg_stat_user_tables
//   pg_stat_user_functions
//   pg_stat_archiver
//   pg_stat_bgwriter
//   pg_stat_slru
//
// You can join recursively with any depth you want (as long as there are no
// cycles), although high depth will incur performance hit. Use at your own risk.
//
package pogo // import "github.com/sanggonlee/pogo"

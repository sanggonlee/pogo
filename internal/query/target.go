package query

import "fmt"

// Target represents a single, non-recursive queryable target.
// It's usually a relation (tables, views) but can be a function result.
type Target int

// Targets defined
const (
	TargetUnspecified Target = iota
	TargetLocks
	TargetLocksOnTxID
	TargetStatActivity
	TargetStatReplication
	TargetStatSSL
	TargetStatGSSAPI
	TargetStatWALReceiver
	TargetStatSubscription
	TargetStatDatabase
	TargetStatDatabaseConflicts
	TargetStatUserTables
	TargetStatUserIndexes
	TargetStatIOUserIndexes
	TargetStatIOUserSequences
	TargetStatIOUserTables
	TargetStatUserFunctions
	TargetStatArchiver
	TargetStatBGWriter
	TargetStatSLRU

	TargetBlockingPIDs

	numTargets
)

// String returns the stringified target.
func (t Target) String() string {
	return [numTargets]string{
		"",
		"pg_locks",
		"pg_locks",
		"pg_stat_activity",
		"pg_stat_replication",
		"pg_stat_ssl",
		"pg_stat_gssapi",
		"pg_stat_wal_receiver",
		"pg_stat_subscription",
		"pg_stat_database",
		"pg_stat_database_conflicts",
		"pg_stat_user_tables",
		"pg_stat_user_indexes",
		"pg_statio_user_indexes",
		"pg_statio_user_sequences",
		"pg_statio_user_tables",
		"pg_stat_user_functions",
		"pg_stat_archiver",
		"pg_stat_bgwriter",
		"pg_stat_slru",

		"",
	}[t]
}

// GetJoinClauses returns the relation alias, column alias, and join condition
// between the target and the target it's being joined with.
func (t Target) GetJoinClauses(join Target) (string, string, string, error) {
	var selectClause, alias, columnAlias, joinCondition string
	var err error

	switch j := (joiner{from: t, join: join}); j {

	// Joined with StatActivity
	case joiner{from: TargetStatActivity, join: TargetLocks}:
		alias = "l"
		columnAlias = "locks"
		joinCondition = "l.pid = pg_stat_activity.pid"
	case joiner{from: TargetStatActivity, join: TargetLocksOnTxID}:
		alias = "txlock"
		columnAlias = "tx_locks"
		joinCondition = "txlock.transactionid = pg_stat_activity.backend_xid"
	case joiner{from: TargetStatActivity, join: TargetStatSSL}:
		alias = "ssl"
		columnAlias = "ssl_usages"
		joinCondition = "ssl.pid = pg_stat_activity.pid"
	case joiner{from: TargetStatActivity, join: TargetStatGSSAPI}:
		alias = "gssapi"
		columnAlias = "gssapi_usages"
		joinCondition = "gssapi.pid = pg_stat_activity.pid"
	case joiner{from: TargetStatActivity, join: TargetStatWALReceiver}:
		alias = "wal_receiver"
		columnAlias = "wal_receivers"
		joinCondition = "wal_receiver.pid = pg_stat_activity.pid"
	case joiner{from: TargetStatActivity, join: TargetStatSubscription}:
		alias = "subscription"
		columnAlias = "subscriptions"
		joinCondition = "subscription.pid = pg_stat_activity.pid"
	case joiner{from: TargetStatActivity, join: TargetStatDatabase}:
		alias = "sa_database"
		columnAlias = "databases"
		joinCondition = "sa_database.datid = pg_stat_activity.datid"
	case joiner{from: TargetStatActivity, join: TargetStatDatabaseConflicts}:
		alias = "database_conflict"
		columnAlias = "database_conflicts"
		joinCondition = "database_conflict.datid = pg_stat_activity.datid"
	case joiner{from: TargetStatActivity, join: TargetBlockingPIDs}:
		return "pg_blocking_pids(pg_stat_activity.pid) AS blocked_by", "", "", nil

	// Joined with StatReplication
	case joiner{from: TargetStatReplication, join: TargetLocks}:
		alias = "l"
		columnAlias = "locks"
		joinCondition = "l.pid = pg_stat_replication.pid"
	case joiner{from: TargetStatReplication, join: TargetStatSSL}:
		alias = "ssl"
		columnAlias = "ssl_usages"
		joinCondition = "ssl.pid = pg_stat_replication.pid"
	case joiner{from: TargetStatReplication, join: TargetStatGSSAPI}:
		alias = "gssapi"
		columnAlias = "gssapi_usages"
		joinCondition = "gssapi.pid = pg_stat_replication.pid"
	case joiner{from: TargetStatReplication, join: TargetStatWALReceiver}:
		alias = "wal_receiver"
		columnAlias = "wal_receivers"
		joinCondition = "wal_receiver.pid = pg_stat_replication.pid"

	// Joined with StatUserTables
	case joiner{from: TargetStatUserTables, join: TargetLocks}:
		alias = "l"
		columnAlias = "locks"
		joinCondition = "l.relation = pg_stat_user_tables.relid"
	case joiner{from: TargetStatUserTables, join: TargetStatUserIndexes}:
		alias = "userindex"
		columnAlias = "indexes"
		joinCondition = "userindex.relid = pg_stat_user_tables.relid"
	case joiner{from: TargetStatUserTables, join: TargetStatSubscription}:
		alias = "subscr"
		columnAlias = "subscriptions"
		joinCondition = "subscr.relid = pg_stat_user_tables.relid"
	case joiner{from: TargetStatUserTables, join: TargetStatIOUserIndexes}:
		alias = "userindex_io"
		columnAlias = "sequence_iostats"
		joinCondition = "userindex_io.relid = pg_stat_user_tables.relid"
	case joiner{from: TargetStatUserTables, join: TargetStatIOUserSequences}:
		alias = "usersequence_io"
		columnAlias = "sequence_iostats"
		joinCondition = "usersequence_io.relid = pg_stat_user_tables.relid"
	case joiner{from: TargetStatUserTables, join: TargetStatIOUserTables}:
		alias = "usertable_io"
		columnAlias = "table_iostats"
		joinCondition = "usertable_io.relid = pg_stat_user_tables.relid"

	// Joined with Locks
	case joiner{from: TargetLocks, join: TargetStatActivity}:
		alias = "locks_sa"
		columnAlias = "activities"
		joinCondition = "locks_sa.pid = pg_locks.pid"
	case joiner{from: TargetLocks, join: TargetStatDatabase}:
		alias = "locks_database"
		columnAlias = "databases"
		joinCondition = "locks_database.datid = pg_locks.database"
	case joiner{from: TargetLocks, join: TargetStatUserTables}:
		alias = "locks_table"
		columnAlias = "tables"
		joinCondition = "locks_table.relid = pg_locks.relation"
	case joiner{from: TargetLocks, join: TargetStatUserIndexes}:
		alias = "locks_index"
		columnAlias = "indexes"
		joinCondition = "locks_index.relid = pg_locks.relation"
	case joiner{from: TargetLocks, join: TargetStatIOUserTables}:
		alias = "locks_table_io"
		columnAlias = "tables_io"
		joinCondition = "locks_table_io.relid = pg_locks.relation"
	case joiner{from: TargetLocks, join: TargetStatIOUserIndexes}:
		alias = "locks_index_io"
		columnAlias = "indexes_io"
		joinCondition = "locks_index_io.relid = pg_locks.relation"
	case joiner{from: TargetLocks, join: TargetStatIOUserSequences}:
		alias = "locks_sequence_io"
		columnAlias = "sequences_io"
		joinCondition = "locks_sequence_io.relid = pg_locks.relation"

	// Joined with StatSSL
	case joiner{from: TargetStatSSL, join: TargetLocks}:
		alias = "ssl_locks"
		columnAlias = "locks"
		joinCondition = "ssl_locks.pid = pg_stat_ssl.pid"
	case joiner{from: TargetStatSSL, join: TargetStatActivity}:
		alias = "ssl_activities"
		columnAlias = "activities"
		joinCondition = "ssl_activities.pid = pg_stat_ssl.pid"

	// Joined with StatGSSAPI
	case joiner{from: TargetStatGSSAPI, join: TargetLocks}:
		alias = "gssapi_locks"
		columnAlias = "locks"
		joinCondition = "gssapi_locks.pid = pg_stat_gssapi.pid"
	case joiner{from: TargetStatGSSAPI, join: TargetStatActivity}:
		alias = "gssapi_activities"
		columnAlias = "activities"
		joinCondition = "gssapi_activities.pid = pg_stat_gssapi.pid"

	// Joined with StatWALReceiver
	case joiner{from: TargetStatWALReceiver, join: TargetLocks}:
		alias = "walreceiver_locks"
		columnAlias = "locks"
		joinCondition = "walreceiver_locks.pid = pg_stat_wal_receiver.pid"
	case joiner{from: TargetStatWALReceiver, join: TargetStatActivity}:
		alias = "walreceiver_activities"
		columnAlias = "activities"
		joinCondition = "walreceiver_activities.pid = pg_stat_wal_receiver.pid"

	// Joined with StatDatabase
	case joiner{from: TargetStatDatabase, join: TargetStatDatabaseConflicts}:
		alias = "statdb_dbconflicts"
		columnAlias = "conflicts"
		joinCondition = "statdb_dbconflicts.datid = pg_stat_database.datid"
	case joiner{from: TargetStatDatabase, join: TargetLocks}:
		alias = "statdb_locks"
		columnAlias = "locks"
		joinCondition = "statdb_locks.database = pg_stat_database.datid"
	case joiner{from: TargetStatDatabase, join: TargetStatActivity}:
		alias = "statdb_activities"
		columnAlias = "activities"
		joinCondition = "statdb_activities.datid = pg_stat_database.datid"

	// Joined with StatSubscription
	case joiner{from: TargetStatSubscription, join: TargetLocks}:
		alias = "subscription_locks"
		columnAlias = "locks"
		joinCondition = "subscription_locks.pid = pg_stat_subscription.pid"
	case joiner{from: TargetStatSubscription, join: TargetStatActivity}:
		alias = "subscription_activities"
		columnAlias = "activities"
		joinCondition = "subscription_activities.pid = pg_stat_subscription.pid"

	// Joined with StatIndex
	case joiner{from: TargetStatUserIndexes, join: TargetStatUserTables}:
		alias = "statindex_tables"
		columnAlias = "tables"
		joinCondition = "statindex_tables.relid = pg_stat_user_indexes.relid"
	case joiner{from: TargetStatUserIndexes, join: TargetStatIOUserTables}:
		alias = "statindex_tablesio"
		columnAlias = "tablesio"
		joinCondition = "statindex_tablesio.relid = pg_stat_user_indexes.relid"
	case joiner{from: TargetStatUserIndexes, join: TargetLocks}:
		alias = "statindex_locks"
		columnAlias = "locks"
		joinCondition = "statindex_locks.relation = pg_stat_user_indexes.indexrelid"
	case joiner{from: TargetStatUserIndexes, join: TargetStatIOUserIndexes}:
		alias = "statindex_indexesio"
		columnAlias = "indexeso"
		joinCondition = "statindex_indexesio.indexrelid = pg_stat_user_indexes.indexrelid"

	default:
		err = j.GetUnsupportedJoinError()
	}

	selectClause = fmt.Sprintf(
		"(CASE WHEN count(%[1]s) = 0 THEN '[]' ELSE json_agg(%[1]s) END) AS %s",
		alias,
		columnAlias,
	)

	return selectClause, alias, joinCondition, err
}

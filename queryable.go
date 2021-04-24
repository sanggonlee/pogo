package pogo

import (
	"github.com/sanggonlee/pogo/internal/query"
	"github.com/sanggonlee/pogo/internal/version"
	"github.com/sanggonlee/pogo/postgres13"
	"github.com/sanggonlee/pogo/postgres9"
)

// Queryable views:
var (
	LocksView                 = query.Queryable{Target: query.TargetLocks}
	LocksOnTxIDView           = query.Queryable{Target: query.TargetLocksOnTxID}
	StatActivityView          = query.Queryable{Target: query.TargetStatActivity}
	StatReplicationView       = query.Queryable{Target: query.TargetStatReplication}
	StatSSLView               = query.Queryable{Target: query.TargetStatSSL}
	StatGSSAPIView            = query.Queryable{Target: query.TargetStatGSSAPI}
	StatWALReceiverView       = query.Queryable{Target: query.TargetStatWALReceiver}
	StatSubscriptionView      = query.Queryable{Target: query.TargetStatSubscription}
	StatDatabaseView          = query.Queryable{Target: query.TargetStatDatabase}
	StatDatabaseConflictsView = query.Queryable{Target: query.TargetStatDatabaseConflicts}
	StatUserTablesView        = query.Queryable{Target: query.TargetStatUserTables}
	StatUserIndexesView       = query.Queryable{Target: query.TargetStatUserIndexes}
	StatIOUserIndexesView     = query.Queryable{Target: query.TargetStatIOUserIndexes}
	StatIOUserSequencesView   = query.Queryable{Target: query.TargetStatIOUserSequences}
	StatIOUserTablesView      = query.Queryable{Target: query.TargetStatIOUserTables}
	StatUserFunctionsView     = query.Queryable{Target: query.TargetStatUserFunctions}
	StatArchiverView          = query.Queryable{Target: query.TargetStatArchiver}
	StatBGWriterView          = query.Queryable{Target: query.TargetStatBGWriter}
	StatSLRUView              = query.Queryable{Target: query.TargetStatSLRU}
)

// Non-relation queryables:
var (
	BlockingPIDs = query.Queryable{Target: query.TargetBlockingPIDs, SelectOnly: true}
)

func setTargets(v version.PostgresVersion) {
	switch v {
	case version.Postgres9:
		StatGSSAPIView.Target = query.TargetUnspecified
		StatSubscriptionView.Target = query.TargetUnspecified
		StatSLRUView.Target = query.TargetUnspecified
	case version.Postgres13:
	}
}

func setSpecifiers(v version.PostgresVersion) {
	switch v {
	case version.Postgres9:
		LocksView.Specifier = &postgres9.Lock{}
		LocksOnTxIDView.Specifier = &postgres9.Lock{}
		StatActivityView.Specifier = &postgres9.StatActivity{}
		StatReplicationView.Specifier = &postgres9.StatReplication{}
		StatSSLView.Specifier = &postgres9.StatSSL{}
		StatGSSAPIView.Specifier = nil // Unsupported
		StatWALReceiverView.Specifier = &postgres9.StatWALReceiver{}
		StatSubscriptionView.Specifier = nil // Unsupported
		StatDatabaseView.Specifier = &postgres9.StatDatabase{}
		StatDatabaseConflictsView.Specifier = &postgres9.StatDatabaseConflict{}
		StatUserTablesView.Specifier = &postgres9.StatTable{}
		StatUserIndexesView.Specifier = &postgres9.StatIndex{}
		StatIOUserIndexesView.Specifier = &postgres9.StatIOIndex{}
		StatIOUserSequencesView.Specifier = &postgres9.StatIOSequence{}
		StatIOUserTablesView.Specifier = &postgres9.StatIOTable{}
		StatUserFunctionsView.Specifier = &postgres9.StatUserFunction{}
		StatArchiverView.Specifier = &postgres9.StatArchiver{}
		StatBGWriterView.Specifier = &postgres9.StatBGWriter{}
		StatSLRUView.Specifier = nil // Unsupported
	case version.Postgres13:
		LocksView.Specifier = &postgres13.Lock{}
		LocksOnTxIDView.Specifier = &postgres13.Lock{}
		StatActivityView.Specifier = &postgres13.StatActivity{}
		StatReplicationView.Specifier = &postgres13.StatReplication{}
		StatSSLView.Specifier = &postgres13.StatSSL{}
		StatGSSAPIView.Specifier = &postgres13.StatGSSAPI{}
		StatWALReceiverView.Specifier = &postgres13.StatWALReceiver{}
		StatSubscriptionView.Specifier = &postgres13.StatSubscription{}
		StatDatabaseView.Specifier = &postgres13.StatDatabase{}
		StatDatabaseConflictsView.Specifier = &postgres13.StatDatabaseConflict{}
		StatUserTablesView.Specifier = &postgres13.StatTable{}
		StatUserIndexesView.Specifier = &postgres13.StatIndex{}
		StatIOUserIndexesView.Specifier = &postgres13.StatIOIndex{}
		StatIOUserSequencesView.Specifier = &postgres13.StatIOSequence{}
		StatIOUserTablesView.Specifier = &postgres13.StatIOTable{}
		StatUserFunctionsView.Specifier = &postgres13.StatUserFunction{}
		StatArchiverView.Specifier = &postgres13.StatArchiver{}
		StatBGWriterView.Specifier = &postgres13.StatBGWriter{}
		StatSLRUView.Specifier = &postgres13.StatSLRU{}
	}
}

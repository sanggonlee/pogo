package postgres9

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatTable represents a row in pg_stat_{all,sys,user}_tables
type StatTable struct {
	RelID                       null.Int    `json:"relid"`
	SchemaName                  null.String `json:"schemaname"`
	RelName                     null.String `json:"relname"`
	NumSequentialScans          null.Int    `json:"seq_scan"`
	NumSequentialRowsRead       null.Int    `json:"seq_tup_read"`
	NumIndexScans               null.Int    `json:"idx_scan"`
	NumIndexRowsFetched         null.Int    `json:"idx_tup_fetch"`
	NumRowsInserted             null.Int    `json:"n_tup_ins"`
	NumRowsUpdated              null.Int    `json:"n_tup_upd"`
	NumRowsDeleted              null.Int    `json:"n_tup_del"`
	NumRowsHotUpdated           null.Int    `json:"n_tup_hot_upd"`
	NumEstimatedLiveRows        null.Int    `json:"n_live_tup"`
	NumEstimatedDeadRows        null.Int    `json:"n_dead_tup"`
	NumRowsModifiedSinceAnalyze null.Int    `json:"n_mod_since_analyze"`
	NumManuallyVacuumed         null.Int    `json:"vacuum_count"`
	LastManuallyVacuumedAt      null.Time   `json:"last_vacuum"`
	NumAutoVacuumed             null.Int    `json:"autovacuum_count"`
	LastAutoVacuumedAt          null.Time   `json:"last_autovacuum"`
	NumManuallyAnalyzed         null.Int    `json:"analyze_count"`
	LastManuallyAnalyzedAt      null.Time   `json:"last_analyze"`
	NumAutoAnalyzed             null.Int    `json:"autoanalyze_count"`
	LastAutoAnalyzedAt          null.Time   `json:"last_autoanalyze"`
}

// Selects returns the column names for select query.
func (s *StatTable) Selects() []string {
	return []string{
		"pg_stat_user_tables.relid",
		"pg_stat_user_tables.schemaname",
		"pg_stat_user_tables.relname",
		"pg_stat_user_tables.seq_scan",
		"pg_stat_user_tables.seq_tup_read",
		"pg_stat_user_tables.idx_scan",
		"pg_stat_user_tables.idx_tup_fetch",
		"pg_stat_user_tables.n_tup_ins",
		"pg_stat_user_tables.n_tup_upd",
		"pg_stat_user_tables.n_tup_del",
		"pg_stat_user_tables.n_tup_hot_upd",
		"pg_stat_user_tables.n_live_tup",
		"pg_stat_user_tables.n_dead_tup",
		"pg_stat_user_tables.n_mod_since_analyze",
		"pg_stat_user_tables.vacuum_count",
		"pg_stat_user_tables.last_vacuum",
		"pg_stat_user_tables.autovacuum_count",
		"pg_stat_user_tables.last_autovacuum",
		"pg_stat_user_tables.analyze_count",
		"pg_stat_user_tables.last_analyze",
		"pg_stat_user_tables.autoanalyze_count",
		"pg_stat_user_tables.last_autoanalyze",
	}
}

// StatTableJoined is the extended struct of StatTable with all the possible joinable fields.
type StatTableJoined struct {
	StatTable
	Locks           Locks           `json:"locks"`
	Indexes         StatIndexes     `json:"indexes"`
	IndexIOStats    StatIndexes     `json:"index_iostats"`
	SequenceIOStats StatIOSequences `json:"sequence_iostats"`
	TableIOStats    StatIOTables    `json:"table_iostats"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatTableJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.RelID,
		&sj.SchemaName,
		&sj.RelName,
		&sj.NumSequentialScans,
		&sj.NumSequentialRowsRead,
		&sj.NumIndexScans,
		&sj.NumIndexRowsFetched,
		&sj.NumRowsInserted,
		&sj.NumRowsUpdated,
		&sj.NumRowsDeleted,
		&sj.NumRowsHotUpdated,
		&sj.NumEstimatedLiveRows,
		&sj.NumEstimatedDeadRows,
		&sj.NumRowsModifiedSinceAnalyze,
		&sj.NumManuallyVacuumed,
		&sj.LastManuallyVacuumedAt,
		&sj.NumAutoVacuumed,
		&sj.LastAutoVacuumedAt,
		&sj.NumManuallyAnalyzed,
		&sj.LastManuallyAnalyzedAt,
		&sj.NumAutoAnalyzed,
		&sj.LastAutoAnalyzedAt,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetStatUserIndexes:
			joinDest = &sj.Indexes
		case query.TargetStatIOUserIndexes:
			joinDest = &sj.IndexIOStats
		case query.TargetStatIOUserSequences:
			joinDest = &sj.SequenceIOStats
		case query.TargetStatIOUserTables:
			joinDest = &sj.TableIOStats
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatTables is an alias for a slice of StatTableJoined.
type StatTables []StatTableJoined

// Scan reads the DB value into StatTables.
func (ss *StatTables) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatTables to a DB value.
func (ss *StatTables) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

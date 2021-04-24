package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatDatabase represents a row in pg_stat_database
type StatDatabase struct {
	DatID               pginternal.OID    `json:"datid,omitempty"`
	DatName             null.String       `json:"datname,omitempty"`
	NumBackends         null.Int          `json:"numbackends,omitempty"`
	XactCommit          pginternal.BigInt `json:"xact_commit,omitempty"`
	XactRollback        pginternal.BigInt `json:"xact_rollback,omitempty"`
	BlocksRead          pginternal.BigInt `json:"blks_read,omitempty"`
	BlocksHit           pginternal.BigInt `json:"blks_hit,omitempty"`
	TuplesReturned      pginternal.BigInt `json:"tup_returned,omitempty"`
	TuplesFetched       pginternal.BigInt `json:"tup_fetched,omitempty"`
	TuplesInserted      pginternal.BigInt `json:"tup_inserted,omitempty"`
	TuplesUpdated       pginternal.BigInt `json:"tup_updated,omitempty"`
	TuplesDeleted       pginternal.BigInt `json:"tup_deleted,omitempty"`
	Conflicts           pginternal.BigInt `json:"conflicts,omitempty"`
	TempFiles           pginternal.BigInt `json:"temp_files,omitempty"`
	TempBytes           pginternal.BigInt `json:"temp_bytes,omitempty"`
	Deadlocks           pginternal.BigInt `json:"deadlocks,omitempty"`
	ChecksumFailures    pginternal.BigInt `json:"checksum_failures,omitempty"`
	ChecksumLastFailure null.Time         `json:"checksum_last_failure,omitempty"`
	BlockReadTime       null.Float        `json:"blk_read_time,omitempty"`
	BlockWriteTime      null.Float        `json:"blk_write_time,omitempty"`
	StatsReset          null.Time         `json:"stats_reset,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatDatabase) Selects() []string {
	return []string{
		"pg_stat_database.datid",
		"pg_stat_database.datname",
		"pg_stat_database.numbackends",
		"pg_stat_database.xact_commit",
		"pg_stat_database.xact_rollback",
		"pg_stat_database.blks_read",
		"pg_stat_database.blks_hit",
		"pg_stat_database.tup_returned",
		"pg_stat_database.tup_fetched",
		"pg_stat_database.tup_inserted",
		"pg_stat_database.tup_updated",
		"pg_stat_database.tup_deleted",
		"pg_stat_database.conflicts",
		"pg_stat_database.temp_files",
		"pg_stat_database.temp_bytes",
		"pg_stat_database.deadlocks",
		"pg_stat_database.checksum_failures",
		"pg_stat_database.checksum_last_failure",
		"pg_stat_database.blk_read_time",
		"pg_stat_database.blk_write_time",
		"pg_stat_database.stats_reset",
	}
}

// StatDatabaseJoined is the extended struct of StatDatabase with all the possible joinable fields.
type StatDatabaseJoined struct {
	StatDatabase

	Conflicts  StatDatabaseConflicts `json:"conflicts"`
	Locks      Locks                 `json:"locks"`
	Activities StatActivities        `json:"activities"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatDatabaseJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.DatID,
		&sj.DatName,
		&sj.NumBackends,
		&sj.XactCommit,
		&sj.XactRollback,
		&sj.BlocksRead,
		&sj.BlocksHit,
		&sj.TuplesReturned,
		&sj.TuplesFetched,
		&sj.TuplesInserted,
		&sj.TuplesUpdated,
		&sj.TuplesDeleted,
		&sj.Conflicts,
		&sj.TempFiles,
		&sj.TempBytes,
		&sj.Deadlocks,
		&sj.ChecksumFailures,
		&sj.ChecksumLastFailure,
		&sj.BlockReadTime,
		&sj.BlockWriteTime,
		&sj.StatsReset,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetStatDatabaseConflicts:
			joinDest = &sj.Conflicts
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetStatActivity:
			joinDest = &sj.Activities
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatDatabases is an alias for a slice of StatDatabaseJoined.
type StatDatabases []StatDatabaseJoined

// Scan reads the DB value into StatDatabases.
func (ss *StatDatabases) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatDatabases to a DB value.
func (ss *StatDatabases) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

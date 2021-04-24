package postgres9

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatIndex represents a row in pg_stat_{all,sys,user}_indexes
type StatIndex struct {
	RelID              pginternal.OID    `json:"relid"`
	IndexRelID         pginternal.OID    `json:"indexrelid"`
	SchemaName         null.String       `json:"schemaname"`
	RelName            null.String       `json:"relname"`
	IndexRelName       null.String       `json:"indexrelname"`
	IndexScan          pginternal.BigInt `json:"idx_scan"`
	IndexTuplesRead    pginternal.BigInt `json:"idx_tup_read"`
	IndexTuplesFetched pginternal.BigInt `json:"idx_tup_fetch"`
}

// Selects returns the column names for select query.
func (s *StatIndex) Selects() []string {
	return []string{
		"pg_stat_user_indexes.relid",
		"pg_stat_user_indexes.indexrelid",
		"pg_stat_user_indexes.schemaname",
		"pg_stat_user_indexes.relname",
		"pg_stat_user_indexes.indexrelname",
		"pg_stat_user_indexes.idx_scan",
		"pg_stat_user_indexes.idx_tup_read",
		"pg_stat_user_indexes.idx_tup_fetch",
	}
}

// StatIndexJoined is the extended struct of StatIndex with all the possible joinable fields.
type StatIndexJoined struct {
	StatIndex

	Tables    StatTables    `json:"tables"`
	TablesIO  StatIOTables  `json:"tables_io"`
	Locks     Locks         `json:"locks"`
	IndexesIO StatIOIndexes `json:"indexes_io"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatIndexJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.RelID,
		&sj.IndexRelID,
		&sj.SchemaName,
		&sj.RelName,
		&sj.IndexRelName,
		&sj.IndexScan,
		&sj.IndexTuplesRead,
		&sj.IndexTuplesFetched,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetStatUserTables:
			joinDest = &sj.Tables
		case query.TargetStatIOUserTables:
			joinDest = &sj.TablesIO
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetStatIOUserIndexes:
			joinDest = &sj.IndexesIO
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatIndexes is an alias for a slice of StatIndexJoined.
type StatIndexes []StatIndexJoined

// Scan reads the DB value into StatIndexes.
func (ss *StatIndexes) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatIndexes to a DB value.
func (ss *StatIndexes) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

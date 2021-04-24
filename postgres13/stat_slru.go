package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatSLRU represents a row in pg_stat_slru view
type StatSLRU struct {
	Name          null.String       `json:"name"`
	BlocksZeroed  pginternal.BigInt `json:"blks_zeroed"`
	BlocksHit     pginternal.BigInt `json:"blks_hit"`
	BlocksRead    pginternal.BigInt `json:"blks_read"`
	BlocksWritten pginternal.BigInt `json:"blks_written"`
	BlocksExists  pginternal.BigInt `json:"blks_exists"`
	Flushes       pginternal.BigInt `json:"flushes"`
	Truncates     pginternal.BigInt `json:"truncates"`
	StatsReset    pginternal.BigInt `json:"stats_reset"`
}

// Selects returns the column names for select query.
func (s *StatSLRU) Selects() []string {
	return []string{
		"pg_stat_slru.name",
		"pg_stat_slru.blks_zeroed",
		"pg_stat_slru.blks_hit",
		"pg_stat_slru.blks_read",
		"pg_stat_slru.blks_written",
		"pg_stat_slru.blks_exists",
		"pg_stat_slru.flushes",
		"pg_stat_slru.truncates",
		"pg_stat_slru.stats_reset",
	}
}

// StatSLRUJoined is the extended struct of StatSLRU with all the possible joinable fields.
type StatSLRUJoined struct {
	StatSLRU
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatSLRUJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.Name,
		&sj.BlocksZeroed,
		&sj.BlocksHit,
		&sj.BlocksRead,
		&sj.BlocksWritten,
		&sj.BlocksExists,
		&sj.Flushes,
		&sj.Truncates,
		&sj.StatsReset,
	}

	return dests
}

// StatSLRUs is an alias for a slice of StatSLRUJoined.
type StatSLRUs []StatSLRUJoined

// Scan reads the DB value into StatSLRUs.
func (ss *StatSLRUs) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatSLRUs to a DB value.
func (ss *StatSLRUs) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

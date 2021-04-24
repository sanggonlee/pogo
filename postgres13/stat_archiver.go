package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatArchiver represents a row in pg_stat_archiver view
type StatArchiver struct {
	ArchivedCount    pginternal.BigInt `json:"archived_count"`
	LastArchivedWAL  null.String       `json:"last_archived_wal"`
	LastArchivedTime null.Time         `json:"last_archived_time"`
	FailedCount      pginternal.BigInt `json:"failed_count"`
	LastFailedWAL    null.String       `json:"last_failed_wal"`
	LastFailedTime   null.Time         `json:"last_failed_time"`
	StatsReset       null.Time         `json:"stats_reset"`
}

// Selects returns the column names for select query.
func (s *StatArchiver) Selects() []string {
	return []string{
		"pg_stat_archiver.archived_count",
		"pg_stat_archiver.last_archived_wal",
		"pg_stat_archiver.last_archived_time",
		"pg_stat_archiver.failed_count",
		"pg_stat_archiver.last_failed_wal",
		"pg_stat_archiver.last_failed_time",
		"pg_stat_archiver.stats_reset",
	}
}

// StatArchiverJoined is the extended struct of StatArchiver with all the possible joinable fields.
type StatArchiverJoined struct {
	StatArchiver
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatArchiverJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.ArchivedCount,
		&sj.LastArchivedWAL,
		&sj.LastArchivedTime,
		&sj.FailedCount,
		&sj.LastFailedWAL,
		&sj.LastFailedTime,
		&sj.StatsReset,
	}

	return dests
}

// StatArchivers is an alias for a slice of StatArchiverJoined.
type StatArchivers []StatArchiverJoined

// Scan reads the DB value into StatArchivers.
func (ss *StatArchivers) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatArchivers to a DB value.
func (ss *StatArchivers) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

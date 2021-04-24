package postgres9

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatBGWriter represents a row in pg_stat_bgwriter view
type StatBGWriter struct {
	CheckpointsTimed    pginternal.BigInt `json:"checkpoints_timed"`
	CheckpointsReq      pginternal.BigInt `json:"checkpoints_req"`
	CheckpointWriteTime null.Float        `json:"checkpoint_write_time"`
	CheckpointSyncTime  null.Float        `json:"checkpoint_sync_time"`
	BuffersCheckpoint   pginternal.BigInt `json:"buffers_checkpoint"`
	BuffersClean        pginternal.BigInt `json:"buffers_clean"`
	MaxWrittenClean     pginternal.BigInt `json:"maxwritten_clean"`
	BuffersBackend      pginternal.BigInt `json:"buffers_backend"`
	BuffersBackendFsync pginternal.BigInt `json:"buffers_backend_fsync"`
	BuffersAlloc        pginternal.BigInt `json:"buffers_alloc"`
	StatsReset          null.Time         `json:"stats_reset"`
}

// Selects returns the column names for select query.
func (s *StatBGWriter) Selects() []string {
	return []string{
		"pg_stat_bgwriter.checkpoints_timed",
		"pg_stat_bgwriter.checkpoints_req",
		"pg_stat_bgwriter.checkpoint_write_time",
		"pg_stat_bgwriter.checkpoint_sync_time",
		"pg_stat_bgwriter.buffers_checkpoint",
		"pg_stat_bgwriter.buffers_clean",
		"pg_stat_bgwriter.maxwritten_clean",
		"pg_stat_bgwriter.buffers_backend",
		"pg_stat_bgwriter.buffers_backend_fsync",
		"pg_stat_bgwriter.buffers_alloc",
		"pg_stat_bgwriter.stats_reset",
	}
}

// StatBGWriterJoined is the extended struct of StatBGWriter with all the possible joinable fields.
type StatBGWriterJoined struct {
	StatBGWriter
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatBGWriterJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.CheckpointsTimed,
		&sj.CheckpointsReq,
		&sj.CheckpointWriteTime,
		&sj.CheckpointSyncTime,
		&sj.BuffersCheckpoint,
		&sj.BuffersClean,
		&sj.MaxWrittenClean,
		&sj.BuffersBackend,
		&sj.BuffersBackendFsync,
		&sj.BuffersAlloc,
		&sj.StatsReset,
	}

	return dests
}

// StatBGWriters is an alias for a slice of StatBGWriterJoined.
type StatBGWriters []StatBGWriterJoined

// Scan reads the DB value into StatBGWriters.
func (ss *StatBGWriters) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatBGWriters to a DB value.
func (ss *StatBGWriters) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

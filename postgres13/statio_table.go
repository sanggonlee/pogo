package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatIOTable represents a row in pg_statio_{all,sys,user}_tables
type StatIOTable struct {
	RelID                pginternal.OID    `json:"relid,omitempty"`
	SchemaName           null.String       `json:"schemaname,omitempty"`
	RelName              null.String       `json:"relname,omitempty"`
	HeapBlocksRead       pginternal.BigInt `json:"heap_blks_read,omitempty"`
	HeapBlocksHit        pginternal.BigInt `json:"heap_blks_hit,omitempty"`
	IndexBlocksRead      pginternal.BigInt `json:"idx_blks_read,omitempty"`
	IndexBlocksHit       pginternal.BigInt `json:"idx_blks_hit,omitempty"`
	ToastBlocksRead      pginternal.BigInt `json:"toast_blks_read,omitempty"`
	ToastBlocksHit       pginternal.BigInt `json:"toast_blks_hit,omitempty"`
	ToastIndexBlocksRead pginternal.BigInt `json:"tidx_blks_read,omitempty"`
	ToastIndexBlocksHit  pginternal.BigInt `json:"tidx_blks_hit,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatIOTable) Selects() []string {
	return []string{
		"pg_statio_user_tables.relid",
		"pg_statio_user_tables.schemaname",
		"pg_statio_user_tables.relname",
		"pg_statio_user_tables.heap_blks_read",
		"pg_statio_user_tables.heap_blks_hit",
		"pg_statio_user_tables.idx_blks_read",
		"pg_statio_user_tables.idx_blks_hit",
		"pg_statio_user_tables.toast_blks_read",
		"pg_statio_user_tables.toast_blks_hit",
		"pg_statio_user_tables.tidx_blks_read",
		"pg_statio_user_tables.tidx_blks_hit",
	}
}

// StatIOTableJoined is the extended struct of StatIOTable with all the possible joinable fields.
type StatIOTableJoined struct {
	StatIOTable
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatIOTableJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.RelID,
		&sj.SchemaName,
		&sj.RelName,
		&sj.HeapBlocksRead,
		&sj.HeapBlocksHit,
		&sj.IndexBlocksRead,
		&sj.IndexBlocksHit,
		&sj.ToastBlocksRead,
		&sj.ToastBlocksHit,
		&sj.ToastIndexBlocksRead,
		&sj.ToastIndexBlocksHit,
	}

	return dests
}

// StatIOTables is an alias for a slice of StatIOTableJoined.
type StatIOTables []StatIOTableJoined

// Scan reads the DB value into StatIOTables.
func (ss *StatIOTables) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatIOTables to a DB value.
func (ss *StatIOTables) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

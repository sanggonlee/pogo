package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatIOIndex represents a row in pg_statio_{all,sys,user}_indexes
type StatIOIndex struct {
	RelID           pginternal.OID    `json:"relid,omitempty"`
	IndexRelID      pginternal.OID    `json:"indexrelid,omitempty"`
	SchemaName      null.String       `json:"schemaname,omitempty"`
	RelName         null.String       `json:"relname,omitempty"`
	IndexRelName    null.String       `json:"indexrelname,omitempty"`
	IndexBlocksRead pginternal.BigInt `json:"idx_blks_read,omitempty"`
	IndexBlocksHit  pginternal.BigInt `json:"idx_blks_hit,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatIOIndex) Selects() []string {
	return []string{
		"pg_statio_user_indexes.relid",
		"pg_statio_user_indexes.indexrelid",
		"pg_statio_user_indexes.schemaname",
		"pg_statio_user_indexes.relname",
		"pg_statio_user_indexes.indexrelname",
		"pg_statio_user_indexes.idx_blks_read",
		"pg_statio_user_indexes.idx_blks_hit",
	}
}

// StatIOIndexJoined is the extended struct of StatIOIndex with all the possible joinable fields.
type StatIOIndexJoined struct {
	StatIOIndex
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatIOIndexJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.RelID,
		&sj.IndexRelID,
		&sj.SchemaName,
		&sj.RelName,
		&sj.IndexRelName,
		&sj.IndexBlocksRead,
		&sj.IndexBlocksHit,
	}

	return dests
}

// StatIOIndexes is an alias for a slice of StatIOIndexJoined.
type StatIOIndexes []StatIOIndexJoined

// Scan reads the DB value into StatIOIndexes.
func (ss *StatIOIndexes) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatIOIndexes to a DB value.
func (ss *StatIOIndexes) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

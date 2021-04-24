package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatIOSequence represents a row in pg_statio_{all,sys,user}_sequences
type StatIOSequence struct {
	RelID      pginternal.OID    `json:"relid,omitempty"`
	SchemaName null.String       `json:"schemaname,omitempty"`
	RelName    null.String       `json:"relname,omitempty"`
	BlocksRead pginternal.BigInt `json:"blks_read,omitempty"`
	BlocksHit  pginternal.BigInt `json:"blks_hit,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatIOSequence) Selects() []string {
	return []string{
		"pg_statio_user_sequences.relid",
		"pg_statio_user_sequences.schemaname",
		"pg_statio_user_sequences.relname",
		"pg_statio_user_sequences.blks_read",
		"pg_statio_user_sequences.blks_hit",
	}
}

// StatIOSequenceJoined is the extended struct of StatIOSequence with all the possible joinable fields.
type StatIOSequenceJoined struct {
	StatIOSequence
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatIOSequenceJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.RelID,
		&sj.SchemaName,
		&sj.RelName,
		&sj.BlocksRead,
		&sj.BlocksHit,
	}

	return dests
}

// StatIOSequences is an alias for a slice of StatIOSequenceJoined.
type StatIOSequences []StatIOSequenceJoined

// Scan reads the DB value into StatIOSequences.
func (ss *StatIOSequences) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatIOSequences to a DB value.
func (ss *StatIOSequences) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

package postgres9

import (
	"database/sql/driver"
	"math/big"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatDatabaseConflict represents a row in pg_stat_database_conflicts
type StatDatabaseConflict struct {
	DatID           pginternal.OID `json:"datid,omitempty"`
	DatName         null.String    `json:"datname,omitempty"`
	ConflTablespace big.Int        `json:"confl_tablespace,omitempty"`
	ConflLock       big.Int        `json:"confl_lock,omitempty"`
	ConflSnapshot   big.Int        `json:"confl_snapshot,omitempty"`
	ConflBufferpin  big.Int        `json:"confl_bufferpin,omitempty"`
	ConflDeadlock   big.Int        `json:"confl_deadlock,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatDatabaseConflict) Selects() []string {
	return []string{
		"pg_stat_database_conflicts.datid",
		"pg_stat_database_conflicts.datname",
		"pg_stat_database_conflicts.confl_tablespace",
		"pg_stat_database_conflicts.confl_lock",
		"pg_stat_database_conflicts.confl_snapshot",
		"pg_stat_database_conflicts.confl_bufferpin",
		"pg_stat_database_conflicts.confl_deadlock",
	}
}

// StatDatabaseConflictJoined is the extended struct of StatDatabaseConflict with all the possible joinable fields.
type StatDatabaseConflictJoined struct {
	StatDatabaseConflict
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatDatabaseConflictJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.DatID,
		&sj.DatName,
		&sj.ConflTablespace,
		&sj.ConflLock,
		&sj.ConflSnapshot,
		&sj.ConflBufferpin,
		&sj.ConflDeadlock,
	}

	return dests
}

// StatDatabaseConflicts is an alias for a slice of StatDatabaseConflictJoined.
type StatDatabaseConflicts []StatDatabaseConflictJoined

// Scan reads the DB value into StatDatabaseConflicts.
func (ss *StatDatabaseConflicts) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatDatabaseConflicts to a DB value.
func (ss *StatDatabaseConflicts) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

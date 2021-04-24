package postgres9

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// Lock represents a row in pg_locks
type Lock struct {
	LockType           null.String `json:"locktype"`
	Database           null.Int    `json:"database"`
	Relation           null.Int    `json:"relation"`
	Page               null.Int    `json:"page"`
	Tuple              null.Int    `json:"tuple"`
	VirtualXID         null.String `json:"virtualxid"`
	TransactionID      null.Int    `json:"transactionid"`
	ClassID            null.Int    `json:"classid"`
	ObjID              null.Int    `json:"objid"`
	ObjSubID           null.Int    `json:"objsubid"`
	VirtualTransaction null.String `json:"virtualtransaction"`
	PID                null.Int    `json:"pid"`
	Mode               null.String `json:"mode"`
	Granted            null.Bool   `json:"granted"`
	FastPath           null.Bool   `json:"fastpath"`
}

// Selects returns the column names for select query.
func (l *Lock) Selects() []string {
	return []string{
		"pg_locks.locktype",
		"pg_locks.database",
		"pg_locks.relation",
		"pg_locks.page",
		"pg_locks.tuple",
		"pg_locks.virtualxid",
		"pg_locks.transactionid",
		"pg_locks.classid",
		"pg_locks.objid",
		"pg_locks.objsubid",
		"pg_locks.virtualtransaction",
		"pg_locks.pid",
		"pg_locks.mode",
		"pg_locks.granted",
		"pg_locks.fastpath",
	}
}

// RowTraceable reports whether the lock has all the information to be able
// to track a specific row in an arbitrary relation.
func (l *Lock) RowTraceable() bool {
	return l.Relation.Valid && l.Page.Valid && l.Tuple.Valid
}

// LockJoined is the extended struct of Lock with all the possible joinable fields.
type LockJoined struct {
	Lock
	Activities  StatActivities  `json:"activities"`
	Databases   StatDatabases   `json:"databases"`
	Tables      StatTables      `json:"tables"`
	Indexes     StatIndexes     `json:"indexes"`
	TablesIO    StatIOTables    `json:"tables_io"`
	IndexesIO   StatIOIndexes   `json:"indexes_io"`
	SequencesIO StatIOSequences `json:"sequences_io"`
	LockedRow   null.String     `json:"locked_row"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (lj *LockJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&lj.LockType,
		&lj.Database,
		&lj.Relation,
		&lj.Page,
		&lj.Tuple,
		&lj.VirtualXID,
		&lj.TransactionID,
		&lj.ClassID,
		&lj.ObjID,
		&lj.ObjSubID,
		&lj.VirtualTransaction,
		&lj.PID,
		&lj.Mode,
		&lj.Granted,
		&lj.FastPath,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetStatActivity:
			joinDest = &lj.Activities
		case query.TargetStatDatabase:
			joinDest = &lj.Databases
		case query.TargetStatUserTables:
			joinDest = &lj.Tables
		case query.TargetStatUserIndexes:
			joinDest = &lj.Indexes
		case query.TargetStatIOUserTables:
			joinDest = &lj.TablesIO
		case query.TargetStatIOUserIndexes:
			joinDest = &lj.IndexesIO
		case query.TargetStatIOUserSequences:
			joinDest = &lj.SequencesIO
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// Locks is an alias for a slice of LockJoined.
type Locks []LockJoined

// Scan reads the DB value into Locks.
func (ls *Locks) Scan(value interface{}) error {
	return convert.JSONScan(ls, value)
}

// Value converts Locks to a DB value.
func (ls *Locks) Value() (driver.Value, error) {
	return convert.JSONValue(ls)
}

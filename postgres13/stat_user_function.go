package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatUserFunction represents a row in pg_stat_user_functions view
type StatUserFunction struct {
	FuncID     pginternal.OID    `json:"funcid"`
	SchemaName null.String       `json:"schemaname"`
	FuncName   null.String       `json:"funcname"`
	Calls      pginternal.BigInt `json:"calls"`
	TotalTime  null.Float        `json:"total_time"`
	SelfTime   null.Float        `json:"self_time"`
}

// Selects returns the column names for select query.
func (s *StatUserFunction) Selects() []string {
	return []string{
		"pg_stat_user_functions.funcid",
		"pg_stat_user_functions.schemaname",
		"pg_stat_user_functions.funcname",
		"pg_stat_user_functions.calls",
		"pg_stat_user_functions.total_time",
		"pg_stat_user_functions.self_time",
	}
}

// StatUserFunctionJoined is the extended struct of StatUserFunction with all the possible joinable fields.
type StatUserFunctionJoined struct {
	StatUserFunction
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatUserFunctionJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.FuncID,
		&sj.SchemaName,
		&sj.FuncName,
		&sj.Calls,
		&sj.TotalTime,
		&sj.SelfTime,
	}

	return dests
}

// StatUserFunctions is an alias for a slice of StatUserFunctionJoined.
type StatUserFunctions []StatUserFunctionJoined

// Scan reads the DB value into StatUserFunctions.
func (ss *StatUserFunctions) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatUserFunctions to a DB value.
func (ss *StatUserFunctions) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

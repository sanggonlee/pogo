package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatGSSAPI represents a row in pg_stat_gssapi
type StatGSSAPI struct {
	PID              null.Int    `json:"pid,omitempty"`
	GSSAuthenticated null.Bool   `json:"gss_authenticated,omitempty"`
	Principal        null.String `json:"principal,omitempty"`
	Encrypted        null.Bool   `json:"encrypted,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatGSSAPI) Selects() []string {
	return []string{
		"pg_stat_gssapi.pid",
		"pg_stat_gssapi.gss_authenticated",
		"pg_stat_gssapi.principal",
		"pg_stat_gssapi.encrypted",
	}
}

// StatGSSAPIJoined is the extended struct of StatGSSAPI with all the possible joinable fields.
type StatGSSAPIJoined struct {
	StatGSSAPI

	Locks      Locks          `json:"locks"`
	Activities StatActivities `json:"activities"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatGSSAPIJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.PID,
		&sj.GSSAuthenticated,
		&sj.Principal,
		&sj.Encrypted,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetStatActivity:
			joinDest = &sj.Activities
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatGSSAPIs is an alias for a slice of StatGSSAPIJoined.
type StatGSSAPIs []StatGSSAPIJoined

// Scan reads the DB value into StatGSSAPIs.
func (ss *StatGSSAPIs) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatGSSAPIs to a DB value.
func (ss *StatGSSAPIs) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

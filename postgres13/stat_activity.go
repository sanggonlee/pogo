package postgres13

import (
	"database/sql/driver"

	"github.com/lib/pq"
	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatActivity represents a row in pg_stat_activity
type StatActivity struct {
	DatID           null.Int    `json:"datid,omitempty"`
	DatName         null.String `json:"datname,omitempty"`
	PID             null.Int    `json:"pid,omitempty"`
	LeaderPID       null.Int    `json:"leader_pid,omitempty"`
	UseSysID        null.Int    `json:"usesysid,omitempty"`
	UseName         null.String `json:"usename,omitempty"`
	ApplicationName null.String `json:"application_name,omitempty"`
	ClientAddr      null.String `json:"client_addr,omitempty"`
	ClientHostname  null.String `json:"client_hostname,omitempty"`
	ClientPort      null.Int    `json:"client_port,omitempty"`
	BackendStart    null.Time   `json:"backend_start,omitempty"`
	XactStart       null.Time   `json:"xact_start,omitempty"`
	QueryStart      null.Time   `json:"query_start,omitempty"`
	StateChange     null.Time   `json:"state_change,omitempty"`
	WaitEventType   null.String `json:"wait_event_type,omitempty"`
	WaitEvent       null.String `json:"wait_event,omitempty"`
	State           null.String `json:"state,omitempty"`
	BackendXID      null.String `json:"backend_xid,omitempty"`
	BackendXMin     null.String `json:"backend_xmin,omitempty"`
	Query           null.String `json:"query,omitempty"`
	BackendType     null.String `json:"backend_type,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatActivity) Selects() []string {
	return []string{
		"pg_stat_activity.datid",
		"pg_stat_activity.datname",
		"pg_stat_activity.pid",
		"pg_stat_activity.leader_pid",
		"pg_stat_activity.usesysid",
		"pg_stat_activity.usename",
		"pg_stat_activity.application_name",
		"pg_stat_activity.client_addr",
		"pg_stat_activity.client_hostname",
		"pg_stat_activity.client_port",
		"pg_stat_activity.backend_start",
		"pg_stat_activity.xact_start",
		"pg_stat_activity.query_start",
		"pg_stat_activity.state_change",
		"pg_stat_activity.wait_event_type",
		"pg_stat_activity.wait_event",
		"pg_stat_activity.state",
		"pg_stat_activity.backend_xid",
		"pg_stat_activity.backend_xmin",
		"pg_stat_activity.query",
		"pg_stat_activity.backend_type",
	}
}

// StatActivityJoined is the extended struct of StatActivity with all the possible joinable fields.
type StatActivityJoined struct {
	StatActivity

	Locks             Locks                 `json:"locks,omitempty"`
	TxLocks           Locks                 `json:"tx_locks,omitempty"`
	SSLUsages         StatSSLs              `json:"ssl_usages,omitempty"`
	GSSAPIUsages      StatGSSAPIs           `json:"gssapi_usages,omitempty"`
	WalRecivers       StatWALReceivers      `json:"wal_receivers,omitempty"`
	Databases         StatDatabases         `json:"databases,omitempty"`
	DatabaseConflicts StatDatabaseConflicts `json:"database_conflicts,omitempty"`
	BlockedBy         pq.Int64Array         `json:"blocked_by,omitempty"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatActivityJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.DatID,
		&sj.DatName,
		&sj.PID,
		&sj.LeaderPID,
		&sj.UseSysID,
		&sj.UseName,
		&sj.ApplicationName,
		&sj.ClientAddr,
		&sj.ClientHostname,
		&sj.ClientPort,
		&sj.BackendStart,
		&sj.XactStart,
		&sj.QueryStart,
		&sj.StateChange,
		&sj.WaitEventType,
		&sj.WaitEvent,
		&sj.State,
		&sj.BackendXID,
		&sj.BackendXMin,
		&sj.Query,
		&sj.BackendType,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetLocksOnTxID:
			joinDest = &sj.TxLocks
		case query.TargetStatSSL:
			joinDest = &sj.SSLUsages
		case query.TargetStatGSSAPI:
			joinDest = &sj.GSSAPIUsages
		case query.TargetStatWALReceiver:
			joinDest = &sj.WalRecivers
		case query.TargetStatDatabase:
			joinDest = &sj.Databases
		case query.TargetBlockingPIDs:
			joinDest = &sj.BlockedBy
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatActivities is an alias for a slice of StatActivityJoined.
type StatActivities []StatActivityJoined

// Scan reads the DB value into StatActivities.
func (ss *StatActivities) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatActivities to a DB value.
func (ss *StatActivities) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

package postgres9

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatReplication represents a row in pg_stat_replication
type StatReplication struct {
	PID             null.Int       `json:"pid,omitempty"`
	UseSysID        null.Int       `json:"usesysid,omitempty"`
	UseName         null.String    `json:"usename,omitempty"`
	ApplicationName null.String    `json:"application_name,omitempty"`
	ClientAddr      null.String    `json:"client_addr,omitempty"`
	ClientHostname  null.String    `json:"client_hostname,omitempty"`
	ClientPort      null.Int       `json:"client_port,omitempty"`
	BackendStart    null.Time      `json:"backend_start,omitempty"`
	BackendXMin     null.String    `json:"backend_xmin,omitempty"`
	State           null.String    `json:"state,omitempty"`
	SentLSN         pginternal.LSN `json:"sent_location,omitempty"`
	WriteLSN        pginternal.LSN `json:"write_location,omitempty"`
	FlushLSN        pginternal.LSN `json:"flush_location,omitempty"`
	ReplayLSN       pginternal.LSN `json:"replay_location,omitempty"`
	SyncPriority    null.Int       `json:"sync_priority,omitempty"`
	SyncState       null.String    `json:"sync_state,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatReplication) Selects() []string {
	return []string{
		"pg_stat_replication.pid",
		"pg_stat_replication.usesysid",
		"pg_stat_replication.usename",
		"pg_stat_replication.application_name",
		"pg_stat_replication.client_addr",
		"pg_stat_replication.client_hostname",
		"pg_stat_replication.client_port",
		"pg_stat_replication.backend_start",
		"pg_stat_replication.backend_xmin",
		"pg_stat_replication.state",
		"pg_stat_replication.sent_location",
		"pg_stat_replication.write_location",
		"pg_stat_replication.flush_location",
		"pg_stat_replication.replay_location",
		"pg_stat_replication.sync_priority",
		"pg_stat_replication.sync_state",
	}
}

// StatReplicationJoined is the extended struct of StatReplication with all the possible joinable fields.
type StatReplicationJoined struct {
	StatReplication

	Locks       Locks            `json:"locks,omitempty"`
	SSLUsages   StatSSLs         `json:"ssl_usages,omitempty"`
	WalRecivers StatWALReceivers `json:"wal_receivers,omitempty"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatReplicationJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.PID,
		&sj.UseSysID,
		&sj.UseName,
		&sj.ApplicationName,
		&sj.ClientAddr,
		&sj.ClientHostname,
		&sj.ClientPort,
		&sj.BackendStart,
		&sj.BackendXMin,
		&sj.State,
		&sj.SentLSN,
		&sj.WriteLSN,
		&sj.FlushLSN,
		&sj.ReplayLSN,
		&sj.SyncPriority,
		&sj.SyncState,
	}

	for _, j := range joins {
		var joinDest interface{}
		switch j.Target {
		case query.TargetLocks:
			joinDest = &sj.Locks
		case query.TargetStatSSL:
			joinDest = &sj.SSLUsages
		case query.TargetStatWALReceiver:
			joinDest = &sj.WalRecivers
		}
		dests = append(dests, joinDest)
	}

	return dests
}

// StatReplications is an alias for a slice of StatReplicationJoined.
type StatReplications []StatReplicationJoined

// Scan reads the DB value into StatReplications.
func (ss *StatReplications) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatReplications to a DB value.
func (ss *StatReplications) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

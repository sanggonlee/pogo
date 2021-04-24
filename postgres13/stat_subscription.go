package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatSubscription represents a row in pg_stat_subscription
type StatSubscription struct {
	SubID              null.Int       `json:"subid,omitempty"`
	SubName            null.String    `json:"subname,omitempty"`
	PID                null.Int       `json:"pid,omitempty"`
	RelID              null.Int       `json:"relid,omitempty"`
	ReceivedLSN        pginternal.LSN `json:"received_lsn,omitempty"`
	LastMsgSendTime    null.Time      `json:"last_msg_send_time,omitempty"`
	LastMsgReceiptTime null.Time      `json:"last_msg_receipt_time,omitempty"`
	LatestEndLSN       pginternal.LSN `json:"latest_end_lsn,omitempty"`
	LatestEndTime      null.Time      `json:"latest_end_time,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatSubscription) Selects() []string {
	return []string{
		"pg_stat_subscription.subid",
		"pg_stat_subscription.subname",
		"pg_stat_subscription.pid",
		"pg_stat_subscription.relid",
		"pg_stat_subscription.received_lsn",
		"pg_stat_subscription.last_msg_send_time",
		"pg_stat_subscription.last_msg_receipt_time",
		"pg_stat_subscription.latest_end_lsn",
		"pg_stat_subscription.latest_end_time",
	}
}

// StatSubscriptionJoined is the extended struct of StatSubscription with all the possible joinable fields.
type StatSubscriptionJoined struct {
	StatSubscription

	Locks      Locks          `json:"locks"`
	Activities StatActivities `json:"activities"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatSubscriptionJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.SubID,
		&sj.SubName,
		&sj.PID,
		&sj.RelID,
		&sj.ReceivedLSN,
		&sj.LastMsgSendTime,
		&sj.LastMsgReceiptTime,
		&sj.LatestEndLSN,
		&sj.LatestEndTime,
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

// StatSubscriptions is an alias for a slice of StatSubscriptionJoined.
type StatSubscriptions []StatSubscriptionJoined

// Scan reads the DB value into StatSubscriptions.
func (ss *StatSubscriptions) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatSubscriptions to a DB value.
func (ss *StatSubscriptions) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/pginternal"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatWALReceiver represents a row in pg_stat_wal_receiver
type StatWALReceiver struct {
	PID                null.Int       `json:"pid,omitempty"`
	Status             null.String    `json:"status,omitempty"`
	ReceiveStartLSN    pginternal.LSN `json:"receive_start_lsn,omitempty"`
	ReceiveStartTLI    null.Int       `json:"receive_start_tli,omitempty"`
	WrittenLSN         pginternal.LSN `json:"written_lsn,omitempty"`
	FlushedLSN         pginternal.LSN `json:"flushed_lsn,omitempty"`
	ReceivedTLI        null.Int       `json:"received_tli,omitempty"`
	LastMsgSendTime    null.Time      `json:"last_msg_send_time,omitempty"`
	LastMsgReceiptTime null.Time      `json:"last_msg_receipt_time,omitempty"`
	LatestEndLSN       pginternal.LSN `json:"latest_end_lsn,omitempty"`
	LatestEndTime      null.Time      `json:"latest_end_time,omitempty"`
	SlotName           null.String    `json:"slot_name,omitempty"`
	SenderHost         null.String    `json:"sender_host,omitempty"`
	SenderPort         null.Int       `json:"sender_port,omitempty"`
	ConnInfo           null.String    `json:"conninfo,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatWALReceiver) Selects() []string {
	return []string{
		"pg_stat_wal_receiver.pid",
		"pg_stat_wal_receiver.status",
		"pg_stat_wal_receiver.receive_start_lsn",
		"pg_stat_wal_receiver.receive_start_tli",
		"pg_stat_wal_receiver.written_lsn",
		"pg_stat_wal_receiver.flushed_lsn",
		"pg_stat_wal_receiver.received_tli",
		"pg_stat_wal_receiver.last_msg_send_time",
		"pg_stat_wal_receiver.last_msg_receipt_time",
		"pg_stat_wal_receiver.latest_end_lsn",
		"pg_stat_wal_receiver.latest_end_time",
		"pg_stat_wal_receiver.slot_name",
		"pg_stat_wal_receiver.sender_host",
		"pg_stat_wal_receiver.sender_port",
		"pg_stat_wal_receiver.conninfo",
	}
}

// StatWALReceiverJoined is the extended struct of StatWALReceiver with all the possible joinable fields.
type StatWALReceiverJoined struct {
	StatWALReceiver

	Locks      Locks          `json:"locks"`
	Activities StatActivities `json:"activities"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatWALReceiverJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.PID,
		&sj.Status,
		&sj.ReceiveStartLSN,
		&sj.ReceiveStartTLI,
		&sj.WrittenLSN,
		&sj.FlushedLSN,
		&sj.ReceivedTLI,
		&sj.LastMsgSendTime,
		&sj.LastMsgReceiptTime,
		&sj.LatestEndLSN,
		&sj.LatestEndTime,
		&sj.SlotName,
		&sj.SenderHost,
		&sj.SenderPort,
		&sj.ConnInfo,
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

// StatWALReceivers is an alias for a slice of StatWALReceiverJoined.
type StatWALReceivers []StatWALReceiverJoined

// Scan reads the DB value into StatWALReceivers.
func (ss *StatWALReceivers) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatWALReceivers to a DB value.
func (ss *StatWALReceivers) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

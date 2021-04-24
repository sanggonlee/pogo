package postgres13

import (
	"database/sql/driver"

	"github.com/sanggonlee/pogo/internal/convert"
	"github.com/sanggonlee/pogo/internal/query"
	"gopkg.in/guregu/null.v3"
)

// StatSSL represents a row in pg_stat_ssl view
type StatSSL struct {
	PID          null.Int    `json:"pid,omitempty"`
	SSL          null.Bool   `json:"ssl,omitempty"`
	Version      null.String `json:"version,omitempty"`
	Cipher       null.String `json:"cipher,omitempty"`
	Bits         null.Int    `json:"bits,omitempty"`
	Compression  null.Bool   `json:"compression,omitempty"`
	ClientDN     null.String `json:"client_dn,omitempty"`
	ClientSerial null.Float  `json:"client_serial,omitempty"`
	IssuerDN     null.String `json:"issuer_dn,omitempty"`
}

// Selects returns the column names for select query.
func (s *StatSSL) Selects() []string {
	return []string{
		"pg_stat_ssl.pid",
		"pg_stat_ssl.ssl",
		"pg_stat_ssl.version",
		"pg_stat_ssl.cipher",
		"pg_stat_ssl.bits",
		"pg_stat_ssl.compression",
		"pg_stat_ssl.client_dn",
		"pg_stat_ssl.client_serial",
		"pg_stat_ssl.issuer_dn",
	}
}

// StatSSLJoined is the extended struct of StatSSL with all the possible joinable fields.
type StatSSLJoined struct {
	StatSSL

	Locks      Locks          `json:"locks"`
	Activities StatActivities `json:"activities"`
}

// ScanDestinations returns the destinations for scanning DB rows in struct fields.
func (sj *StatSSLJoined) ScanDestinations(joins []query.Queryable) []interface{} {
	dests := []interface{}{
		&sj.PID,
		&sj.SSL,
		&sj.Version,
		&sj.Cipher,
		&sj.Bits,
		&sj.Compression,
		&sj.ClientDN,
		&sj.ClientSerial,
		&sj.IssuerDN,
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

// StatSSLs is an alias for a slice of StatSSLJoined.
type StatSSLs []StatSSLJoined

// Scan reads the DB value into StatSSLs.
func (ss *StatSSLs) Scan(value interface{}) error {
	return convert.JSONScan(ss, value)
}

// Value converts StatSSLs to a DB value.
func (ss *StatSSLs) Value() (driver.Value, error) {
	return convert.JSONValue(ss)
}

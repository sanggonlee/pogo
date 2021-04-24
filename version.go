package pogo

import (
	"errors"
	"fmt"

	"github.com/sanggonlee/pogo/internal/version"
)

// PostgresVersion represents the version of Postgres pogo will use.
// Currently supports 9.6 and 13
type PostgresVersion version.PostgresVersion

// String representation of Postgres version.
func (v PostgresVersion) String() string {
	return []string{
		"9.6",
		"13",
	}[v]
}

// Postgres version enums
const (
	Postgres9  PostgresVersion = PostgresVersion(version.Postgres9)
	Postgres13                 = PostgresVersion(version.Postgres13)
)

var (
	errVersionAlreadySet = errors.New("postgres version already set")
)

var defaultPostgresVersion = Postgres13

// SetPostgresVersion sets the Postgres version and locks it down.
// You can set the version only once. Trying to set it again will return a
// non-nil error and won't have any effects.
func SetPostgresVersion(v PostgresVersion) error {
	if version.IsSet() {
		return errVersionAlreadySet
	}

	_version := version.PostgresVersion(v)
	version.Set(_version)
	setTargets(_version)
	setSpecifiers(_version)

	return nil
}

// GetPostgresVersion reports the current Postgres version pogo is using.
func GetPostgresVersion() PostgresVersion {
	return PostgresVersion(version.Get())
}

// getVersionMismatchError returns an error for when version doesn't match between
// the operation requested and the currently set version.
func getVersionMismatchError(asked, current PostgresVersion) error {
	return fmt.Errorf("unable to run version %s specific action for current version %s", asked, current)
}

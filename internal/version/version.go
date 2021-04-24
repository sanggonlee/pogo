package version

// PostgresVersion represents the current Postgres version pogo is targeting for.
type PostgresVersion int

// Supports 9.6 and 13.
const (
	Postgres9 PostgresVersion = iota
	Postgres13
)

var isSet bool
var _version PostgresVersion

// IsSet reports whether the version was previously set.
func IsSet() bool {
	return isSet
}

// Set sets the Postgres version pogo is using.
func Set(v PostgresVersion) {
	_version = v
	isSet = true
}

// Get returns the Postgres version set.
func Get() PostgresVersion {
	return _version
}

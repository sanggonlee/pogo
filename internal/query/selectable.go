package query

// Selectable is the interface describing the relations
// a select query can be run for.
type Selectable interface {
	// Selects returns a list of column names as appeared
	// in a select query.
	Selects() []string
}

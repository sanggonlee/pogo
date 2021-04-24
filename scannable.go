package pogo

import "github.com/sanggonlee/pogo/internal/query"

// Scannable describes objects that qualify for Postgres query destinations
type Scannable interface {
	// ScanDestinations returns all the scan destinations of the
	// receiver, along with the destinations of all the joined relations,
	// which are given as its arguments.
	ScanDestinations([]query.Queryable) []interface{}
}

package query

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	errMutualRecursionDetected = errors.New("mutual recursion detected")
)

// Queryable represents an abstract entity that you can run query against.
type Queryable struct {
	Target     Target
	Specifier  Selectable
	Joins      []Queryable
	SelectOnly bool

	where string
}

// With appends "child" Queryables to the target Queryable which will be included
// in the resulting parent data object.
func (q Queryable) With(queryables ...Queryable) Queryable {
	_q := q
	_q.Joins = append(q.Joins, queryables...)
	return _q
}

// Where specifies the sql WHERE clause for the queryable.
func (q Queryable) Where(where string) Queryable {
	_q := q
	_q.where = where
	return _q
}

// ToQuery converts the Queryable to a SQL query.
func (q Queryable) ToQuery() (string, error) {
	// Prevent infinite recursion
	if q.mutualRecursionDetected() {
		return "", errMutualRecursionDetected
	}

	return q.toQuery()
}

func (q Queryable) getUnsupportedTargetError() error {
	return fmt.Errorf("queryable %s is not supported in this version", q)
}

func (q Queryable) toQuery() (string, error) {
	if q.Target == TargetUnspecified || (!q.SelectOnly && q.Specifier == nil) {
		return "", q.getUnsupportedTargetError()
	}

	var selects []string
	if q.Specifier != nil {
		selects = q.Specifier.Selects()
	}
	var groupBys []string
	if len(q.Joins) > 0 {
		groupBys = selects
	}
	var where string
	if q.where != "" {
		where = fmt.Sprintf("WHERE %s", q.where)
	}

	joins := make([]string, 0, len(q.Joins))
	for _, j := range q.Joins {
		joinQuery, err := j.toQuery()
		if err != nil {
			return "", errors.Wrapf(err, "converting join query for %s under %s", j.Target, q.Target)
		}

		columnFromJoin, joinAlias, joinCondition, err := q.Target.GetJoinClauses(j.Target)
		if err != nil {
			return "", errors.Wrap(err, "getting join clauses")
		}

		selects = append(selects, columnFromJoin)
		if !j.SelectOnly {
			joins = append(joins, fmt.Sprintf(`LEFT JOIN (
				%s
			) AS %s ON %s`, joinQuery, joinAlias, joinCondition))
		}
	}

	var groupBy string
	if len(groupBys) > 0 {
		groupBy = fmt.Sprintf("GROUP BY %s", strings.Join(groupBys, ", "))
	}

	return fmt.Sprintf(`SELECT
		%s
		FROM %s
		%s
		%s
		%s`,
		strings.Join(selects, ", "),
		q.Target,
		strings.Join(joins, ", "),
		where,
		groupBy,
	), nil
}

// String is a string representation of Queryable
func (q Queryable) String() string {
	joins := make([]string, 0, len(q.Joins))
	for _, j := range q.Joins {
		joins = append(joins, j.String())
	}
	return fmt.Sprintf(
		`Queryable{target: %s, specifier: %s, selectOnly: %t, joins: %s}`,
		q.Target,
		q.Specifier,
		q.SelectOnly,
		strings.Join(joins, ", "),
	)
}

func (q Queryable) mutualRecursionDetected() bool {
	m := make(map[string]bool)
	return searchForSameQueryable(q, m)
}

func searchForSameQueryable(q Queryable, seen map[string]bool) bool {
	if seen[q.String()] {
		return true
	}
	seen[q.String()] = true

	for _, j := range q.Joins {
		if searchForSameQueryable(j, seen) {
			return true
		}
	}

	return false
}

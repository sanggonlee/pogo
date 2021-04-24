package query

import (
	"fmt"
)

type joiner struct {
	from Target
	join Target
}

func (j joiner) GetUnsupportedJoinError() error {
	return fmt.Errorf(
		"join between %s and %s is not supported",
		j.from,
		j.join,
	)
}

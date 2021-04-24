package convert

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONScan scans a db value as a JSON.
func JSONScan(dest, value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.([]byte); ok {
		if err := json.Unmarshal(v, dest); err != nil {
			return err
		}
	}

	return nil
}

// JSONValue converts a JSON value into a db value.
func JSONValue(src interface{}) (driver.Value, error) {
	if src == nil {
		return nil, nil
	}

	v, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	return v, nil
}

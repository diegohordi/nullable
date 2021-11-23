package nullable

import (
	"bytes"
	"database/sql"
	"time"
)

type Time struct {
	sql.NullTime
}

func NewTime(v time.Time) *Time {
	return &Time{sql.NullTime{
		Time:  v,
		Valid: true,
	}}
}

func (n Time) MarshalJSON() ([]byte, error) {
	if !n.Valid || n.Time.IsZero() {
		return jsonNullBytes, nil
	}
	return n.Time.MarshalJSON()
}

func (n *Time) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) || bytes.Equal(data, jsonEmptyBytes) {
		n.Valid = false
		return nil
	}
	err := n.Time.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

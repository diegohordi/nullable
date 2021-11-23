package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Int64 struct {
	sql.NullInt64
}

func NewInt64(v int64) *Int64 {
	return &Int64{sql.NullInt64{
		Int64: v,
		Valid: true,
	}}
}

func (n Int64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.Int64)
}

func (n *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.Int64)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

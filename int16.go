package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Int16 struct {
	sql.NullInt16
}

func NewInt16(v int16) *Int16 {
	return &Int16{sql.NullInt16{
		Int16: v,
		Valid: true,
	}}
}

func (n Int16) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.Int16)
}

func (n *Int16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.Int16)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Int32 struct {
	sql.NullInt32
}

func NewInt32(v int32) *Int32 {
	return &Int32{sql.NullInt32{
		Int32: v,
		Valid: true,
	}}
}

func (n Int32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.Int32)
}

func (n *Int32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.Int32)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Float64 struct {
	sql.NullFloat64
}

func NewFloat64(v float64) *Float64 {
	return &Float64{sql.NullFloat64{
		Float64: v,
		Valid:   true,
	}}
}

func (n Float64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.Float64)
}

func (n *Float64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.Float64)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type Bool struct {
	sql.NullBool
}

func NewBool(v bool) *Bool {
	return &Bool{sql.NullBool{
		Bool:  v,
		Valid: true,
	}}
}

func (n Bool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.Bool)
}

func (n *Bool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.Bool)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

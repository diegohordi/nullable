package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func NewString(v string) *String {
	return &String{sql.NullString{
		String: v,
		Valid:  true,
	}}
}

func (n String) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return jsonNullBytes, nil
	}
	return json.Marshal(n.String)
}

func (n *String) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullBytes) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(data, &n.String)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

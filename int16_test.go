package nullable_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/diegohordi/nullable"
	"math"
	"reflect"
	"testing"
)

func TestInt16_MarshalJSON(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "should return null",
			fields: fields{
				value: nullable.Int16{
					NullInt16: sql.NullInt16{
						Int16: 0,
						Valid: false,
					},
				}},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "should return the given int16",
			fields: fields{
				value: *nullable.NewInt16(math.MaxInt16),
			},
			want:    []byte(fmt.Sprintf("%v", math.MaxInt16)),
			wantErr: false,
		},
		{
			name: "should marshal the given int16 from a struct",
			fields: fields{
				value: &struct {
					ID    int            `json:"id"`
					Value nullable.Int16 `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewInt16(100),
				},
			},
			want:    []byte(`{"id":100,"value":100}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.fields.value)
			if err != nil && tt.wantErr {
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		holder  interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name: "should unmarshal a null value",
			args: args{
				data: []byte("null"),
			},
			holder: &nullable.Int16{},
			want: &nullable.Int16{
				NullInt16: sql.NullInt16{
					Int16: 0,
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should unmarshal into a struct",
			args: args{
				data: []byte(`{"id": 100, "value":120}`),
			},
			holder: &struct {
				ID    int            `json:"id"`
				Value nullable.Int16 `json:"value"`
			}{},
			want: &struct {
				ID    int            `json:"id"`
				Value nullable.Int16 `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewInt16(120),
			},
			wantErr: false,
		},
		{
			name: "should return an error due to an unexpected value",
			args: args{
				data: []byte(`{"id": 100, "value":"test"}`),
			},
			holder: &struct {
				ID    int            `json:"id"`
				Value nullable.Int16 `json:"value"`
			}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal(tt.args.data, tt.holder)
			if err != nil && tt.wantErr {
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.holder, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", tt.holder, tt.want)
			}
		})
	}
}

func TestInt16_Scan(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    nullable.Int16
		wantErr bool
	}{
		{
			name: "should return a nullable int16 with 0 as value",
			fields: fields{
				value: nil,
			},
			want: nullable.Int16{
				NullInt16: sql.NullInt16{
					Int16: 0,
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should return a nullable int16 with the given value as its value",
			fields: fields{
				value: 100,
			},
			want:    *nullable.NewInt16(100),
			wantErr: false,
		},
		{
			name: "should return an error due to an unsupported format",
			fields: fields{
				value: make(chan string),
			},
			want:    nullable.Int16{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n nullable.Int16
			err := n.Scan(tt.fields.value)
			if err != nil && tt.wantErr {
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("Scan() got = %v, want %v", n, tt.want)
			}
		})
	}
}

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

func TestInt64_MarshalJSON(t *testing.T) {
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
				value: nullable.Int64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
				}},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "should return the given int64",
			fields: fields{
				value: *nullable.NewInt64(math.MaxInt64),
			},
			want:    []byte(fmt.Sprintf("%v", math.MaxInt64)),
			wantErr: false,
		},
		{
			name: "should marshal the given int64 from a struct",
			fields: fields{
				value: &struct {
					ID    int            `json:"id"`
					Value nullable.Int64 `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewInt64(100),
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

func TestInt64_UnmarshalJSON(t *testing.T) {
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
			holder: &nullable.Int64{},
			want: &nullable.Int64{
				NullInt64: sql.NullInt64{
					Int64: 0,
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
				Value nullable.Int64 `json:"value"`
			}{},
			want: &struct {
				ID    int            `json:"id"`
				Value nullable.Int64 `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewInt64(120),
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
				Value nullable.Int64 `json:"value"`
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

func TestInt64_Scan(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    nullable.Int64
		wantErr bool
	}{
		{
			name: "should return a nullable int64 with 0 as value",
			fields: fields{
				value: nil,
			},
			want: nullable.Int64{
				NullInt64: sql.NullInt64{
					Int64: 0,
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should return a nullable int64 with the given value as its value",
			fields: fields{
				value: 100,
			},
			want:    *nullable.NewInt64(100),
			wantErr: false,
		},
		{
			name: "should return an error due to an unsupported format",
			fields: fields{
				value: make(chan string),
			},
			want:    nullable.Int64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n nullable.Int64
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

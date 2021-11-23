package nullable_test

import (
	"database/sql"
	"encoding/json"
	"github.com/diegohordi/nullable"
	"reflect"
	"testing"
)

func TestBool_MarshalJSON(t *testing.T) {
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
				value: nullable.Bool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: false,
					},
				}},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "should return the given true boolean",
			fields: fields{
				value: *nullable.NewBool(true),
			},
			want:    []byte(`true`),
			wantErr: false,
		},
		{
			name: "should return the given false boolean",
			fields: fields{
				value: *nullable.NewBool(false),
			},
			want:    []byte(`false`),
			wantErr: false,
		},
		{
			name: "should marshal the given true boolean from a struct",
			fields: fields{
				value: &struct {
					ID    int           `json:"id"`
					Value nullable.Bool `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewBool(true),
				},
			},
			want:    []byte(`{"id":100,"value":true}`),
			wantErr: false,
		},
		{
			name: "should marshal the given false boolean from a struct",
			fields: fields{
				value: &struct {
					ID    int           `json:"id"`
					Value nullable.Bool `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewBool(false),
				},
			},
			want:    []byte(`{"id":100,"value":false}`),
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

func TestBool_UnmarshalJSON(t *testing.T) {
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
			holder: &nullable.Bool{},
			want: &nullable.Bool{
				NullBool: sql.NullBool{
					Bool:  false,
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should unmarshal a true boolean into a struct",
			args: args{
				data: []byte(`{"id": 100, "value":true}`),
			},
			holder: &struct {
				ID    int           `json:"id"`
				Value nullable.Bool `json:"value"`
			}{},
			want: &struct {
				ID    int           `json:"id"`
				Value nullable.Bool `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewBool(true),
			},
			wantErr: false,
		},
		{
			name: "should unmarshal a false boolean into a struct",
			args: args{
				data: []byte(`{"id": 100, "value":false}`),
			},
			holder: &struct {
				ID    int           `json:"id"`
				Value nullable.Bool `json:"value"`
			}{},
			want: &struct {
				ID    int           `json:"id"`
				Value nullable.Bool `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewBool(false),
			},
			wantErr: false,
		},
		{
			name: "should return an error due to an unexpected value",
			args: args{
				data: []byte(`{"id": 100, "value":"test"}`),
			},
			holder: &struct {
				ID    int           `json:"id"`
				Value nullable.Bool `json:"value"`
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

func TestBool_Scan(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    nullable.Bool
		wantErr bool
	}{
		{
			name: "should return a nullable bool with false as value",
			fields: fields{
				value: nil,
			},
			want: nullable.Bool{
				NullBool: sql.NullBool{
					Bool:  false,
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should return a nullable bool with the given value (true) as its value",
			fields: fields{
				value: true,
			},
			want:    *nullable.NewBool(true),
			wantErr: false,
		},
		{
			name: "should return a nullable bool with the given value (false) as its value",
			fields: fields{
				value: false,
			},
			want:    *nullable.NewBool(false),
			wantErr: false,
		},
		{
			name: "should return an error due to an unsupported format",
			fields: fields{
				value: make(chan string),
			},
			want:    nullable.Bool{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n nullable.Bool
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

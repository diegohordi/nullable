package nullable_test

import (
	"database/sql"
	"encoding/json"
	"github.com/diegohordi/nullable"
	"reflect"
	"testing"
)

func TestString_MarshalJSON(t *testing.T) {
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
				value: nullable.String{
					NullString: sql.NullString{
						String: "",
						Valid:  false,
					},
				}},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "should return an empty string",
			fields: fields{
				value: *nullable.NewString(""),
			},
			want:    []byte(`""`),
			wantErr: false,
		},
		{
			name: "should return the given string",
			fields: fields{
				value: *nullable.NewString("test"),
			},
			want:    []byte(`"test"`),
			wantErr: false,
		},
		{
			name: "should marshal the given string from a struct",
			fields: fields{
				value: &struct {
					ID    int             `json:"id"`
					Value nullable.String `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewString("test"),
				},
			},
			want:    []byte(`{"id":100,"value":"test"}`),
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

func TestString_UnmarshalJSON(t *testing.T) {
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
			name: "should unmarshal a empty string",
			args: args{
				data: []byte(`""`),
			},
			holder:  &nullable.String{},
			want:    nullable.NewString(""),
			wantErr: false,
		},
		{
			name: "should unmarshal a null string",
			args: args{
				data: []byte("null"),
			},
			holder: &nullable.String{},
			want: &nullable.String{
				NullString: sql.NullString{
					String: "",
					Valid:  false,
				}},
			wantErr: false,
		},
		{
			name: "should unmarshal into a struct",
			args: args{
				data: []byte(`{"id": 100, "value":"test"}`),
			},
			holder: &struct {
				ID    int             `json:"id"`
				Value nullable.String `json:"value"`
			}{},
			want: &struct {
				ID    int             `json:"id"`
				Value nullable.String `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewString("test"),
			},
			wantErr: false,
		},
		{
			name: "should return an error due to an unexpected value",
			args: args{
				data: []byte(`{"id": 100, "value":false}`),
			},
			holder: &struct {
				ID    int             `json:"id"`
				Value nullable.String `json:"value"`
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

func TestString_Scan(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    nullable.String
		wantErr bool
	}{
		{
			name: "should return a nullable string with an empty string as value",
			fields: fields{
				value: nil,
			},
			want: nullable.String{
				NullString: sql.NullString{
					String: "",
					Valid:  false,
				}},
			wantErr: false,
		},
		{
			name: "should return a nullable string with the given value as its value",
			fields: fields{
				value: "test",
			},
			want:    *nullable.NewString("test"),
			wantErr: false,
		},
		{
			name: "should return an error due to an unsupported format",
			fields: fields{
				value: make(chan string),
			},
			want:    nullable.String{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n nullable.String
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

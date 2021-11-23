package nullable_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/diegohordi/nullable"
	"time"

	"reflect"
	"testing"
)

var (
	timeRef    = time.Date(2021, 11, 23, 12, 10, 0, 0, time.UTC)
	timeRefStr = "2021-11-23T12:10:00Z"
)

func TestTime_MarshalJSON(t *testing.T) {
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
				value: *nullable.NewTime(time.Time{}),
			},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "should return the given time",
			fields: fields{
				value: *nullable.NewTime(timeRef),
			},
			want:    []byte(fmt.Sprintf(`"%s"`, timeRefStr)),
			wantErr: false,
		},
		{
			name: "should marshal the given time from a struct",
			fields: fields{
				value: &struct {
					ID    int           `json:"id"`
					Value nullable.Time `json:"value"`
				}{
					ID:    100,
					Value: *nullable.NewTime(timeRef),
				},
			},
			want:    []byte(fmt.Sprintf(`{"id":100,"value":"%s"}`, timeRefStr)),
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

func TestTime_UnmarshalJSON(t *testing.T) {
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
			holder: &nullable.Time{},
			want: &nullable.Time{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should unmarshal a null value",
			args: args{
				data: []byte("null"),
			},
			holder: &nullable.Time{},
			want: &nullable.Time{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				}},
			wantErr: false,
		},
		{
			name: "should unmarshal into a struct",
			args: args{
				data: []byte(fmt.Sprintf(`{"id": 100, "value":"%s"}`, timeRefStr)),
			},
			holder: &struct {
				ID    int           `json:"id"`
				Value nullable.Time `json:"value"`
			}{},
			want: &struct {
				ID    int           `json:"id"`
				Value nullable.Time `json:"value"`
			}{
				ID:    100,
				Value: *nullable.NewTime(timeRef),
			},
			wantErr: false,
		},
		{
			name: "should return an error due to an unexpected value",
			args: args{
				data: []byte(`{"id": 100, "value":false}`),
			},
			holder: &struct {
				ID    int           `json:"id"`
				Value nullable.Time `json:"value"`
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

func TestTime_Scan(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    nullable.Time
		wantErr bool
	}{
		{
			name: "should return a nullable time with a zero time as value",
			fields: fields{
				value: nil,
			},
			want: nullable.Time{NullTime: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			}},
			wantErr: false,
		},
		{
			name: "should return a nullable time with the given value as its value",
			fields: fields{
				value: timeRef,
			},
			want:    *nullable.NewTime(timeRef),
			wantErr: false,
		},
		{
			name: "should return an error due to an unsupported format",
			fields: fields{
				value: make(chan string),
			},
			want:    nullable.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n nullable.Time
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

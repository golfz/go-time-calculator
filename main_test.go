package main

import (
	"reflect"
	"testing"
)

func TestExtractTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want MyTime
	}{
		{
			name: "empty string, expect all 0",
			args: args{s: ""},
			want: MyTime{
				Hour:   0,
				Minute: 0,
				Second: 0,
			},
		},
		{
			name: "only second, expect 0,0,s",
			args: args{s: "09"},
			want: MyTime{
				Hour:   0,
				Minute: 0,
				Second: 9,
			},
		},
		{
			name: "minute, second, expect 0,m,s",
			args: args{s: "02:09"},
			want: MyTime{
				Hour:   0,
				Minute: 2,
				Second: 9,
			},
		},
		{
			name: "hour, minute, second, expect h,m,s",
			args: args{s: "11:02:09"},
			want: MyTime{
				Hour:   11,
				Minute: 2,
				Second: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractTime(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

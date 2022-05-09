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

func TestAddTime(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all 00:00:00, expect 00:00:00",
			args: args{
				a: "00:00:00",
				b: "00:00:00",
			},
			want: "00:00:00",
		},
		{
			name: "no overflow, expect 12:12:12",
			args: args{
				a: "02:04:07",
				b: "10:08:05",
			},
			want: "12:12:12",
		},
		{
			name: "second overflow, expect 12:12:12",
			args: args{
				a: "02:03:17",
				b: "10:08:55",
			},
			want: "12:12:12",
		},
		{
			name: "minute overflow, expect 12:12:12",
			args: args{
				a: "01:14:07",
				b: "10:58:05",
			},
			want: "12:12:12",
		},
		{
			name: "minute, second overflow, expect 12:12:12",
			args: args{
				a: "01:13:17",
				b: "10:58:55",
			},
			want: "12:12:12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTime(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

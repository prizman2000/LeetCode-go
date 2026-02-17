package stringToIntegerAtoi

import (
	"reflect"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{s: "+42"}, want: 42},
		{name: "Example 2", args: args{s: " -042"}, want: -42},
		{name: "Example 2", args: args{s: "1337c0d3"}, want: 1337},
		{name: "Example 1", args: args{s: "0-1"}, want: 0},
		{name: "Example 2", args: args{s: "words and 987"}, want: 0},
		{name: "Example 2", args: args{s: "-92147483648"}, want: -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

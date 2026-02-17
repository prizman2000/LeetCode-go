package zigzagConversion

import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 1", args: args{s: "PAYPALISHIRING", i: 3}, want: "PAHNAPLSIIGYIR"},
		{name: "Example 2", args: args{s: "PAYPALISHIRING", i: 4}, want: "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

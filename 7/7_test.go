package reverseInteger

import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{i: 123}, want: 321},
		{name: "Example 2", args: args{i: 100}, want: 1},
		{name: "Example 2", args: args{i: 22022}, want: 22022},
		{name: "Example 1", args: args{i: -123}, want: -321},
		{name: "Example 2", args: args{i: -100}, want: -1},
		{name: "Example 2", args: args{i: -37022022}, want: -22022073},
		{name: "Example 2", args: args{i: 1534236469}, want: 0},
		{name: "Example 2", args: args{i: -1534236469}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

package longestPalindromicSubstr

import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 1", args: args{s: "babad"}, want: "bab"},
		{name: "Example 2", args: args{s: "cbbd"}, want: "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

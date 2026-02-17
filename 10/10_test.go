package regularExpressionMatching

import (
	"reflect"
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s, p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{s: "qwe", p: "qwe"}, want: true},
		{name: "Example 2", args: args{s: "qwe", p: "ewq"}, want: false},
		{name: "Example 3", args: args{s: "qre", p: "q.e"}, want: true},
		{name: "Example 4", args: args{s: "eew", p: "..e"}, want: false},
		{name: "Example 5", args: args{s: "eew", p: "e*w"}, want: true},
		{name: "Example 6", args: args{s: "qweqweqwe", p: "q.e..eq*w*e"}, want: true},
		{name: "Example 7", args: args{s: "qweqweqwe", p: "q..*.e"}, want: true},
		{name: "Example 8", args: args{s: "aa", p: "a"}, want: false},
		{name: "Example 9", args: args{s: "aab", p: "c*a*b"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

package palindromeNumber

import (
	"reflect"
	"testing"
)

func Test_palindromeNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{x: -123321}, want: false},
		{name: "Example 2", args: args{x: 123321}, want: true},
		{name: "Example 3", args: args{x: 12321}, want: true},
		{name: "Example 4", args: args{x: 1}, want: true},
		{name: "Example 5", args: args{x: 0}, want: true},
		{name: "Example 6", args: args{x: 235675}, want: false},
		{name: "Example 7", args: args{x: 55}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

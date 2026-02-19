package threeSum

import (
	"reflect"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 2",
			args: args{nums: []int{0, 1, 1}},
			want: [][]int{},
		},
		{
			name: "Example 3",
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Example 4",
			args: args{nums: []int{-1, 0, 1, 0}},
			want: [][]int{{1, 0, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

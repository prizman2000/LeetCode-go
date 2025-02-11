package removeDuplicatesFromSortedArray

import (
	"reflect"
	"testing"
)

func Test_removeDuplicatesFromSortedArray(t *testing.T) {
	type args struct {
		sortedArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{sortedArray: []int{1, 1, 2}}, want: 2},
		{name: "Example 2", args: args{sortedArray: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.sortedArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

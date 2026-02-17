package maxArea

import (
	"testing"
)

func Test_maxAres(t *testing.T) {
	type args struct {
		height []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "Example 2",
			args: args{height: []int{1, 1}},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{height: []int{4, 3, 2, 1, 4}},
			want: 16,
		},
		{
			name: "Example 4",
			args: args{height: []int{0, 2}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

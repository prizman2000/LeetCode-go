package longestCommonPrefix

import (
	"reflect"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "Example 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "Example 3",
			args: args{strs: []string{"ab", "a"}},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

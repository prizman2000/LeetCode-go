package intToRoman

import (
	"reflect"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{num: 3749},
			want: "MMMDCCXLIX",
		},
		{
			name: "Example 2",
			args: args{num: 58},
			want: "LVIII",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

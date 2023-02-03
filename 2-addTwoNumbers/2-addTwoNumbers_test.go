package addTwoNumbers

import (
	"reflect"
	"testing"
)

func makeList(input []int) *ListNode {

	var head *ListNode = nil
	var next *ListNode = head

	for _, digit := range input {
		next = &ListNode{digit, nil}
		next = next.Next
	}

	return head
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{makeList([]int{2, 4, 3}), makeList([]int{5, 6, 4})}, makeList([]int{7, 0, 8})},
		{"Example 2", args{makeList([]int{0}), makeList([]int{0})}, makeList([]int{0})},
		{"Example 3", args{makeList([]int{9, 9, 9, 9, 9, 9, 9}), makeList([]int{9, 9, 9, 9})}, makeList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

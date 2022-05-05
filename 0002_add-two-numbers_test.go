package leetcode_go

import (
	"reflect"
	"testing"
)

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
		{
			name: "example1",
			args: args{
				l1: newListNode(2, 4, 3),
				l2: newListNode(5, 6, 4),
			},
			want: newListNode(7, 0, 8),
		},
		{
			name: "example2",
			args: args{
				l1: newListNode(0),
				l2: newListNode(0),
			},
			want: newListNode(0),
		},
		{
			name: "example3",
			args: args{
				l1: newListNode(9, 9, 9, 9, 9, 9, 9),
				l2: newListNode(9, 9, 9, 9),
			},
			want: newListNode(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newListNode(inputs ...int) *ListNode {
	begin := &ListNode{Val: inputs[0]}
	cursor := begin
	for _, v := range inputs[1:] {
		cursor.Next = &ListNode{Val: v}
		cursor = cursor.Next
	}

	return begin
}

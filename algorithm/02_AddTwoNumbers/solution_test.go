package _02_AddTwoNumbers

import (
	"fmt"
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
			name: "right-1",
			args: args{
				l1: toListNode([]int{2, 4, 3}),
				l2: toListNode([]int{5, 6, 4}),
			},
			want: toListNode([]int{7, 0, 8}),
		},
		{
			name: "right-2",
			args: args{
				l1: toListNode([]int{2, 4, 5}),
				l2: toListNode([]int{5, 6, 4}),
			},
			want: toListNode([]int{7, 0, 0, 1}),
		},
		{
			name: "right-3",
			args: args{
				l1: toListNode([]int{1}),
				l2: toListNode([]int{9, 9, 9, 9}),
			},
			want: toListNode([]int{0, 0, 0, 0, 1}),
		},
		{
			name: "right-4",
			args: args{
				l1: toListNode([]int{0}),
				l2: toListNode([]int{9, 9, 9, 9}),
			},
			want: toListNode([]int{9, 9, 9, 9}),
		},
		{
			name: "right-5",
			args: args{
				l1: toListNode([]int{}),
				l2: toListNode([]int{9, 9, 9, 9}),
			},
			want: toListNode([]int{9, 9, 9, 9}),
		},
		{
			name: "right-6",
			args: args{
				l1: toListNode([]int{}),
				l2: toListNode([]int{}),
			},
			want: toListNode([]int{0}),
		},
		{
			name: "right-7",
			args: args{
				l1: toListNode([]int{0}),
				l2: toListNode([]int{0}),
			},
			want: toListNode([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !equal(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got.toString(), tt.want.toString())
			}
		})
	}
}

func equal(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}

func (l *ListNode) toString() string {
	if l == nil {
		return "0"
	}
	s := fmt.Sprintf("%d", l.Val)
	for l.Next != nil {
		s += fmt.Sprintf("->%d", l.Next.Val)
		l = l.Next
	}
	return s
}

func toListNode(nums []int) *ListNode {
	result := &ListNode{}
	current := result
	for _, v := range nums {
		current.Next = &ListNode{
			Val: v,
		}
		current = current.Next
	}
	return result.Next
}

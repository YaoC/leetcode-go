package _02_AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/submissions/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	previous := result
	v1, v2, carry := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			v1 = 0
		} else {
			v1, l1 = l1.Val, l1.Next
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2, l2 = l2.Val, l2.Next
		}
		current.Val, carry = (v1+v2+carry)%10, (v1+v2+carry)/10
		current.Next = &ListNode{}
		previous = current
		current = current.Next
	}
	if carry == 1 {
		current.Val = 1
	}
	if current.Val == 0 {
		previous.Next = nil
	}
	return result
}

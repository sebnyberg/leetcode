package p0083dedupsortedlist

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
		Val:  math.MinInt16,
	}
	cur := dummy
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next, cur = next, next
	}
	return dummy.Next
}

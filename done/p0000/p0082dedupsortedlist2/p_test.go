package p0082dedupsortedlist2

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	head := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							Next: &ListNode{
								Val: 5,
							},
							Val: 4,
						},
						Val: 4,
					},
					Val: 3,
				},
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}

	deleteDuplicates(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Next: head,
		Val:  -101,
	}
	cur := dummy
	for cur.Next != nil {
		next := cur.Next
		if next.Next != nil && next.Next.Val == next.Val {
			nextVal := next.Val
			for next != nil && next.Val == nextVal {
				next = next.Next
			}
			cur.Next = next
			continue
		}
		cur.Next, cur = next, next
	}
	return dummy.Next
}

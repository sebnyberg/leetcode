package p0143reorderlist

import "testing"

func Test_reorderList(t *testing.T) {
	head := &ListNode{Val: 1}
	cur := head
	n := 7
	for i := 0; i < n; i++ {
		cur.Next = new(ListNode)
		cur.Next.Val = i + 2
		cur = cur.Next
	}

	reorderList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow is now just before the middle of the list
	// start reversing all pointers in second section of the list
	cur, next := slow, slow.Next
	for next != nil {
		next.Next, next, cur = cur, next.Next, next
	}

	// mid is now at the end
	// start weaving pointers until the nodes meet in the middle
	start := head
	end := cur

	for {
		if start == end {
			start.Next = nil
			return
		} else if end.Next == start {
			start.Next, end.Next = end, nil
			return
		}
		start, start.Next, end, end.Next = start.Next, end, end.Next, start.Next
	}
}

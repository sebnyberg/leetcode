package p0092reverselinkedlist2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	// Find the last node prior to the segment to reverse
	preLeft := dummy
	for i := 1; i < left; i++ {
		preLeft = preLeft.Next
	}

	// Reverse pointers within segment
	first := preLeft.Next
	cur := first
	next := cur.Next
	for i := left; i < right; i++ {
		cur, next, next.Next = next, next.Next, cur
	}

	// Point last node prior to the segment to the last element of the segment
	preLeft.Next = cur

	// Point the first element beyond the segment to the first element
	// within the segment.
	first.Next = next

	return dummy.Next
}

package p2816doubleanumberrepresentedasalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	// Iterate over the linked list and reverse edges along the way.
	var prev *ListNode
	curr := head
	next := head.Next
	head.Next = nil
	for next != nil {
		// Starts with
		// prev <- current | next
		//
		// We shift to
		// prev <- current | next
		//         current <- next | next.Next
		//          (prev) 	(curr)  (next)
		prev, curr, next, next.Next = curr, next, next.Next, curr
	}

	// Iterate over the linked list, reversing edges in exactly the same way once
	// again, but this time, multiply by two, add the carry, and store carry for
	// next round.
	var carry int
	next = prev
	curr.Next = nil
	curr.Val = (curr.Val + carry) * 2
	carry = curr.Val / 10
	curr.Val %= 10

	for next != nil {
		prev, curr, next, next.Next = curr, next, next.Next, curr
		curr.Val = (curr.Val * 2) + carry
		carry = curr.Val / 10
		curr.Val %= 10
	}

	// Handle the head node
	if carry != 0 {
		head = &ListNode{Val: carry, Next: head}
	}
	return head
}

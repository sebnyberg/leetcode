package p0061rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 1
	last := head
	for ; last.Next != nil; last = last.Next {
		n++
	}
	last.Next = head
	k %= n
	last = head
	for i := 0; i < n-k-1; i++ {
		last = last.Next
	}
	head = last.Next
	last.Next = nil
	return head
}

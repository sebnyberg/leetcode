package p0142linkedlistcycle2

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

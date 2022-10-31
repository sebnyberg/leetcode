package p0876middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy
	for {
		slow = slow.Next
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	return slow
}

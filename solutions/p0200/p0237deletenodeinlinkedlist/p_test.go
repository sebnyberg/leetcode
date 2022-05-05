package p0237deletenodeinlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	cur, next := node, node.Next
	for {
		cur.Val = next.Val
		if next.Next == nil {
			cur.Next = nil
			return
		}
		cur, next = cur.Next, next.Next
	}
}

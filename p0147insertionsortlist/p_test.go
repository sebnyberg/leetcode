package p0147insertionsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)

	cur := head
	for cur != nil {
		next := cur.Next
		insertNode := dummy
		for insertNode.Next != nil && insertNode.Next.Val <= cur.Val {
			insertNode = insertNode.Next
		}
		insertNode.Next, cur.Next = cur, insertNode.Next

		// do next
		cur = next
	}
	return dummy.Next
}

package p0141linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	smallNumber := -100001
	cur := head
	for cur != nil {
		if cur.Val == smallNumber {
			return true
		}
		cur.Val = smallNumber
		cur = cur.Next
	}
	return false
}

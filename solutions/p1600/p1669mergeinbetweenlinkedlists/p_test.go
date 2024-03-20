package p1669mergeinbetweenlinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var pre, post *ListNode
	curr := list1
	for i := 0; i < a-1; i++ {
		curr = curr.Next
	}
	pre = curr
	for i := 0; i < b-a+2; i++ {
		curr = curr.Next
	}
	post = curr
	pre.Next = list2
	curr = list2
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = post
	return list1
}

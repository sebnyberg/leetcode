package p2807insertgreatestcommondivicorsinlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var gcd func(m, n int) int
	gcd = func(m, n int) int {
		if n == 0 {
			return m
		}
		return gcd(n, m%n)
	}
	curr := head
	for curr.Next != nil {
		curr.Next = &ListNode{
			Val:  gcd(curr.Val, curr.Next.Val),
			Next: curr.Next,
		}
		curr = curr.Next.Next
	}
	return head
}

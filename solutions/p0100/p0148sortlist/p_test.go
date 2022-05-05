package p0148sortlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortList(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.in, tc.want), func(t *testing.T) {
			root := &ListNode{Val: tc.in[0]}
			cur := root
			for _, v := range tc.in[1:] {
				cur.Next = &ListNode{
					Val: v,
				}
				cur = cur.Next
			}
			cur = sortList(root)
			for _, v := range tc.want {
				require.Equal(t, v, cur.Val)
				cur = cur.Next
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// A single node is always sorted
	if head == nil || head.Next == nil {
		return head
	}

	// Split list in two using a slow and fast iterator
	// Yields lists of even size, or k/2 and k/2+1
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// Capture head of second section, then cut off first from second
	secondHead := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(secondHead))
}

// Merge two sorted lists
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// Merge into l1 from l2
	// Ensure that l1 holds the starting value:
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	head := l1
	for l2 != nil { // While there is stuff to merge form l2
		// Move l1 forward when it is smaller than l2
		for l1.Next != nil && l1.Next.Val < l2.Val {
			l1 = l1.Next
		}
		// Put one item from l2 into l1
		l2, l2.Next, l1.Next = l2.Next, l1.Next, l2
	}
	return head
}

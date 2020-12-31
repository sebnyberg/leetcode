package l0019_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeNthFromEnd(t *testing.T) {
	for _, tc := range []struct {
		values []int
		n      int
		want   *ListNode
	}{
		{[]int{1, 2, 3, 4, 5}, 2, createList([]int{1, 2, 3, 5})},
		{[]int{1}, 1, createList([]int{})},
		{[]int{1, 2}, 1, createList([]int{1})},
		{[]int{1, 2}, 2, createList([]int{2})},
	} {
		t.Run(fmt.Sprintf("%+v", tc.values), func(t *testing.T) {
			root := createList(tc.values)
			require.Equal(t, tc.want, removeNthFromEnd(root, tc.n))
		})
	}
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	root := &ListNode{
		Val: values[0],
	}
	prev := root
	for i := 1; i < len(values); i++ {
		prev.Next = &ListNode{
			Val: values[i],
		}
		prev = prev.Next
	}
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// For n=1, do the easy solution
	if n == 1 {
		if head.Next == nil {
			return nil
		}
		prev := head
		for {
			if prev.Next.Next == nil {
				prev.Next = nil
				return head
			}
			prev = prev.Next
		}
	}

	// Once the end is reached, the n+1th, n-1th nodes need to be available
	// to remove the nth node, i.e. n+1 total nodes need to be in the stack
	nodeStack := make([]*ListNode, n+1)
	nodeStack[0] = head

	next := head.Next
	for i := 1; i < n+1; i++ {
		if next == nil {
			return head.Next
		}
		nodeStack[i] = next
		next = next.Next
	}

	// Move stack forward until the end
	for nodeStack[n].Next != nil {
		for i := range nodeStack {
			nodeStack[i] = nodeStack[i].Next
		}
	}

	// Remove element
	nodeStack[0].Next = nodeStack[2]

	return head
}

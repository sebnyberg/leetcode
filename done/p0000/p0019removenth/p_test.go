package p0019removenth

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
	if head == nil {
		return head
	}

	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		// n == len(nodes)
		// return head.Next
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

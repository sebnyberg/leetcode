package d24_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n.Val))
	for cur := n.Next; cur != nil; cur = cur.Next {
		sb.WriteString(" > ")
		sb.WriteString(strconv.Itoa(cur.Val))
	}
	return sb.String()
}

func createList(in []int) *ListNode {
	var prev *ListNode
	for i := len(in) - 1; i >= 0; i-- {
		prev = &ListNode{
			Val:  in[i],
			Next: prev,
		}
	}
	return prev
}

func Test_createList(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want string
	}{
		{[]int{1, 2, 3, 4}, "1 > 2 > 3 > 4"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, createList(tc.in).String())
		})
	}
}

func Test_swapPairs(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want string
	}{
		{[]int{1, 2}, "2 > 1"},
		{[]int{1, 2, 3}, "2 > 1 > 3"},
		{[]int{1, 2, 3, 4}, "2 > 1 > 4 > 3"},
		{[]int{1}, "1"},
		// {[]int{1, 2, 3, 4}, "1 > 2 > 3 > 4"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, swapPairs(createList(tc.in)).String())
		})
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	// (1) -> (2) -> Maybe(3)
	// Swap to (2) -> (1) -> Maybe(3)
	head, head.Next, head.Next.Next = head.Next, head.Next.Next, head

	for cur, prev := head.Next.Next, head.Next; cur != nil && cur.Next != nil; {
		prev.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		prev, cur = cur, cur.Next
	}

	return head
}

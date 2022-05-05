package p1721swappingnodesinlinkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_swapNodes(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{100, 90}, 2, []int{90, 100}},
		{[]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}, 5, []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}},
		{[]int{1, 2, 3}, 2, []int{1, 2, 3}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 4, 3, 2, 5}},
	} {
		t.Run(fmt.Sprintf("%+v/%v/%+v", tc.in, tc.k, tc.want), func(t *testing.T) {
			root := &ListNode{Val: tc.in[0]}
			cur := root
			for _, v := range tc.in[1:] {
				cur.Next = &ListNode{
					Val: v,
				}
				cur = cur.Next
			}
			cur = swapNodes(root, tc.k)
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

func swapNodes(head *ListNode, k int) *ListNode {
	first := head
	for k > 1 {
		k--
		first = first.Next
	}
	second := head
	lead := first
	for lead.Next != nil {
		lead = lead.Next
		second = second.Next
	}

	first.Val, second.Val = second.Val, first.Val
	return head
}

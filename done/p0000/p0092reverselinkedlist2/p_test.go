package p0092reverselinkedlist2

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

func createList(values []int) *ListNode {
	nodes := make([]*ListNode, len(values))
	for i := len(values) - 1; i >= 0; i-- {
		nodes[i] = &ListNode{
			Val: values[i],
		}
		if i < len(values)-1 {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes[0]
}

func (h *ListNode) String() string {
	cur := h
	nodeVals := make([]string, 0)
	for cur != nil {
		nodeVals = append(nodeVals, strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	return strings.Join(nodeVals, ",")
}

func Test_partition(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		m    int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 4, []int{4, 3, 2, 1, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
	} {
		t.Run(fmt.Sprintf("%+v/%v/%v", tc.in, tc.m, tc.n), func(t *testing.T) {
			input := createList(tc.in)
			wantHead := createList(tc.want)
			got := reverseBetween(input, tc.m, tc.n)
			require.Equal(t, wantHead.String(), got.String())
		})
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	before := dummy
	for i := 0; i < m-1; i++ {
		before = before.Next
		if before == nil {
			return dummy.Next
		}
	}

	if before.Next == nil || before.Next.Next == nil {
		return dummy.Next
	}

	// Reverse the nested list
	prev, cur, first := before.Next, before.Next.Next, before.Next
	for i := 0; i < n-m; i++ {
		prev, cur, cur.Next = cur, cur.Next, prev
		if cur == nil {
			break
		}
	}
	before.Next, first.Next = prev, cur
	return dummy.Next
}

package p0086partitionlist

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
		x    int
		want []int
	}{
		{[]int{1, 4, 3, 0, 2, 5, 2}, 3, []int{1, 0, 2, 2, 4, 3, 5}},
		// {[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		// {[]int{2, 1}, 2, []int{1, 2}},
		// {[]int{1}, 2, []int{1}},
		// {[]int{3}, 2, []int{3}},
		// {[]int{4, 5, 6}, 2, []int{4, 5, 6}},
		// {[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.in, tc.x), func(t *testing.T) {
			input := createList(tc.in)
			wantHead := createList(tc.want)
			got := partition(input, tc.x)
			require.Equal(t, wantHead.String(), got.String())
		})
	}
}

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{
		Next: head,
		Val:  -101,
	}

	// Find the first node that is greater than or equal to the target
	// This node will serve as the previous node for all nodes below
	// the target further into the list
	first := dummy
	for first.Next != nil && first.Next.Val < x {
		first = first.Next
	}

	// At this point we know that first.Next is part of the second partition
	if first.Next == nil {
		return dummy.Next
	}
	second := first.Next

	for second != nil && second.Next != nil {
		if second.Next.Val < x {
			first, first.Next, second.Next.Next, second.Next = second.Next, second.Next, first.Next, second.Next.Next
			continue
		}
		second = second.Next
	}

	return dummy.Next
}

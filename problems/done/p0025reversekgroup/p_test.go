package p0025reversekgroup

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

func (l *ListNode) String() string {
	if l == nil {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	cur := l
	for cur.Next != nil {
		sb.WriteString(strconv.Itoa(cur.Val) + ",")
		cur = cur.Next
	}
	sb.WriteString(strconv.Itoa(cur.Val))
	sb.WriteString("]")
	return sb.String()
}

func Test_reverseKGroup(t *testing.T) {
	for _, tc := range []struct {
		head []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.head, tc.k), func(t *testing.T) {
			require.Equal(t, createList(tc.want), reverseKGroup(createList(tc.head), tc.k))
		})
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	for {
		// Check that len(remainder) > k
		for cur, i := prev, 0; i < k; i++ {
			if cur.Next == nil {
				return dummy.Next
			}
			cur = cur.Next
		}
		// Reverse section
		cur, nextPrev := prev.Next, prev.Next
		for i := 0; i < k; i++ {
			prev, cur, cur.Next = cur, cur.Next, prev
		}
		nextPrev.Next, nextPrev.Next.Next, prev = cur, prev, nextPrev
	}
}

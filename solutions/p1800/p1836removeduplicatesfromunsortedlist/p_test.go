package p1836removeduplicatesfromunsortedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicatesUnsorted(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{2, 1, 1, 2}, []int{}},
		{[]int{1, 2, 3, 2}, []int{1, 3}},
		{[]int{3, 2, 2, 1, 3, 2, 4}, []int{1, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			testList := &ListNode{
				Val: tc.nums[0],
			}
			cur := testList
			n := len(tc.nums)
			for pos := 1; pos < n; pos++ {
				cur.Next = &ListNode{
					Val: tc.nums[pos],
				}
				cur = cur.Next
			}
			res := deleteDuplicatesUnsorted(testList)
			cur = res
			pos := 0
			for cur != nil {
				require.True(t, pos < len(tc.want), "invalid length")
				require.Equal(t, tc.want[pos], cur.Val)
				pos++
				cur = cur.Next
			}
			require.Equal(t, pos, len(tc.want))
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	// Count how many times each value has been seen
	// Map of bool serves as a three-state counter
	// Does not exist -> false -> true (second count onwards)
	seen := make(map[int]bool)
	cur := head
	for cur != nil {
		if _, exists := seen[cur.Val]; !exists {
			seen[cur.Val] = false
		} else {
			seen[cur.Val] = true
		}
		cur = cur.Next
	}
	// Remove duplicate values
	cur = dummy
	for cur.Next != nil {
		if seen[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

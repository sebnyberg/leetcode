package p0021mergesortedlists

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func createList(values []int) *ListNode {
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	root := &ListNode{
// 		Val: values[0],
// 	}
// 	prev := root
// 	for i := 1; i < len(values); i++ {
// 		prev.Next = &ListNode{
// 			Val: values[i],
// 		}
// 		prev = prev.Next
// 	}
// 	return root
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func Test_mergeTwoLists(t *testing.T) {
// 	for _, tc := range []struct {
// 		l1   []int
// 		l2   []int
// 		want []int
// 	}{
// 		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
// 		{[]int{}, []int{}, []int{}},
// 		{[]int{}, []int{0}, []int{0}},
// 	} {
// 		t.Run(fmt.Sprintf("%+v/+%v", tc.l1, tc.l2), func(t *testing.T) {
// 			require.Equal(t, createList(tc.want), mergeTwoLists(createList(tc.l1), createList(tc.l2)))
// 		})
// 	}
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	for {
// 		if l1 == nil {
// 			return l2
// 		}
// 		if l2 == nil {
// 			return l1
// 		}

// 	}
// }

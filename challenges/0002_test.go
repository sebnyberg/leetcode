package leetcode_test

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

func Test_l0002(t *testing.T) {
	tcs := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v,%v", tc.l1, tc.l2), func(t *testing.T) {
			l1 := l0002_createlistNode(tc.l1)
			l2 := l0002_createlistNode(tc.l2)
			require.EqualValues(t, tc.want, l0002_addTwoNumbers(l1, l2).Ints())
		})
	}
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	cur := new(ListNode)
// 	res := cur

// 	// Create copy of l1 inside res
// 	for {
// 		cur.Val = l1.Val
// 		cur.Next = l1.Next
// 		if l1.Next == nil {
// 			break
// 		}
// 		cur = l1.Next
// 		l1 = l1.Next
// 	}

// 	// Add l2 values to res
// 	var nilNode = &ListNode{Val: 0}
// 	var carry int
// 	var sum int
// 	cur = res
// 	for {
// 		sum = cur.Val + l2.Val + carry
// 		carry = sum / 10
// 		cur.Val = sum % 10

// 		if cur.Next == nil && l2.Next == nil {
// 			if carry == 1 {
// 				cur.Next = &ListNode{Val: carry}
// 			}
// 			break
// 		}

// 		// When there are no more values in l2, refer to a nilNode
// 		if l2.Next == nil {
// 			l2 = nilNode
// 		} else {
// 			l2 = l2.Next
// 		}

// 		if cur.Next == nil {
// 			cur.Next = &ListNode{Val: 0}
// 		}
// 		cur = cur.Next
// 	}
// 	return res
// }

// v2 re-uses l1 to generate output
func l0002_addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := l1

	// Add l2 values to res
	var nilNode = &ListNode{Val: 0}
	var carry int
	var sum int

	for {
		sum = l1.Val + l2.Val + carry
		carry = sum / 10
		l1.Val = sum % 10

		if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
				l1.Next = &ListNode{Val: carry}
			}
			break
		}

		// When there are no more values in l2, refer to a nilNode
		if l2.Next == nil {
			l2 = nilNode
		} else {
			l2 = l2.Next
		}

		if l1.Next == nil {
			l1.Next = &ListNode{Val: 0}
		}
		l1 = l1.Next
	}
	return res
}

func (l *ListNode) Ints() []int {
	res := make([]int, 1)
	res[0] = l.Val
	for cur := l; cur.Next != nil; cur = cur.Next {
		res = append(res, cur.Next.Val)
	}
	return res
}

func (l *ListNode) String() string {
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

func l0002_createlistNode(vals []int) *ListNode {
	first := &ListNode{
		Val: vals[0],
	}
	cur := first
	for i := 1; i < len(vals); i++ {
		next := ListNode{
			Val: vals[i],
		}
		if i == 1 {
			first.Next = &next
		} else {
			cur.Next = &next
		}
		cur = &next
	}
	return first
}

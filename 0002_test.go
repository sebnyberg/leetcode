package leetcode_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_addTwo(t *testing.T) {
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
			l1 := createListNode(tc.l1)
			l2 := createListNode(tc.l2)
			res := addTwoNumbers(l1, l2)
			fmt.Println(res)
		})
	}
	t.FailNow()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	res := cur

	// Create copy of l1 inside res
	for {
		cur.Val = l1.Val
		cur.Next = l1.Next
		if l1.Next == nil {
			break
		}
		cur = l1.Next
		l1 = l1.Next
	}

	// Add l2 values to res
	var nilNode = &ListNode{Val: 0}
	var carry int
	var sum int
	cur = res
	for {
		sum = cur.Val + l2.Val + carry
		carry = sum / 10
		cur.Val = sum % 10

		if cur.Next == nil && l2.Next == nil {
			if carry == 1 {
				cur.Next = &ListNode{Val: carry}
			}
			break
		}

		// When there are no more values in l2, refer to a nilNode
		if l2.Next == nil {
			l2 = nilNode
		} else {
			l2 = l2.Next
		}

		if cur.Next == nil {
			cur.Next = &ListNode{Val: 0}
		}
		cur = cur.Next
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

func createListNode(vals []int) *ListNode {
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

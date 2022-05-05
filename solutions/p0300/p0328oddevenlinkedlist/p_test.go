package p0328oddevenlinkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOddEvenLinkedList(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 3, 5, 7, 2, 4, 6, 8}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			dummy := &ListNode{}
			cur := dummy
			for _, num := range tc.input {
				cur.Next = &ListNode{
					Val: num,
				}
				cur = cur.Next
			}
			res := oddEvenList(dummy.Next)
			cur = res
			got := []int{}
			for cur != nil {
				got = append(got, cur.Val)
				cur = cur.Next
			}
			require.Equal(t, tc.want, got)
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

func printList(head *ListNode) {
	i := 0
	for head != nil {
		if i > 0 {
			fmt.Print("->")
		}
		fmt.Print(head.Val)
		head = head.Next
		i++
	}
	fmt.Print("\n")
}

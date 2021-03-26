package p0234palindromelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	n := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	require.Equal(t, true, isPalindrome(n))
	n = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	require.Equal(t, true, isPalindrome(n))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	stack := make([]int, 0)
	for {
		stack = append(stack, slow.Val)
		if fast.Next == nil {
			stack = stack[:len(stack)-1]
			break
		} else if fast.Next.Next == nil {
			break
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	slow = slow.Next
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return true
}

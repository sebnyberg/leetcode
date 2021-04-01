package p0234palindromelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	n := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	require.Equal(t, true, isPalindrome(n))
	n = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	require.Equal(t, true, isPalindrome(n))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode
	rev, next := head, head.Next
	rev.Next = prev
	for next != slow {
		prev, rev, next = rev, next, next.Next
		rev.Next = prev
	}

	if fast != nil { // odd
		slow = slow.Next
	}

	// Compare
	for {
		switch {
		case slow == nil && rev == nil:
			return true
		case slow == nil && rev != nil,
			rev == nil && slow != nil:
			return false
		case slow.Val != rev.Val:
			return false
		}
		slow, rev = slow.Next, rev.Next
	}
}

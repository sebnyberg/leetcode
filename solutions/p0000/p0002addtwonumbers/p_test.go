package p0002addtwonumbers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ParseList(s string) *ListNode {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	items := strings.Split(s, ",")
	dummy := &ListNode{}
	cur := dummy
	for _, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("failed to convert string %v to integer", num)
		}
		cur.Next = &ListNode{
			Val: num,
		}
		cur = cur.Next
	}
	return dummy.Next
}

func (n *ListNode) String() string {
	cur := n
	var sb strings.Builder
	sb.WriteRune('[')
	i := 0
	for cur != nil {
		if i != 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(cur.Val))
		i++
		cur = cur.Next
	}
	sb.WriteRune(']')
	return sb.String()
}

func (n *ListNode) Equals(other *ListNode) bool {
	if n == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	return n.Val == other.Val && n.Next.Equals(other.Next)
}

func Test_addTwoNumbers(t *testing.T) {
	tcs := []struct {
		l1   string
		l2   string
		want string
	}{
		{"[2,4,3]", "[5,6,4]", "[7,0,8]"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v,%v", tc.l1, tc.l2), func(t *testing.T) {
			res := addTwoNumbers(ParseList(tc.l1), ParseList(tc.l2))
			require.True(t, ParseList(tc.want).Equals(res))
		})
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	val := func(l *ListNode) int {
		if l != nil {
			return l.Val
		}
		return 0
	}
	next := func(l *ListNode) *ListNode {
		if l != nil {
			return l.Next
		}
		return l
	}
	curr := dummy
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = val(l1) + val(l2) + carry
		carry = curr.Val / 10
		curr.Val %= 10
		l1 = next(l1)
		l2 = next(l2)
	}
	return dummy.Next
}

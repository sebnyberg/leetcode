package p2181mergenodesinbetweenzeros

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_mergeNodes(t *testing.T) {
	for _, tc := range []struct {
		head string
		want string
	}{
		{"[0,3,1,0,4,5,2,0]", "[4,11]"},
		{"[0,1,0,3,0,2,2,0]", "[1,3,4]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			head := ParseList(tc.head)
			res := mergeNodes(head)
			require.Equal(t, tc.want, res.String())
		})
	}
}

func mergeNodes(head *ListNode) *ListNode {
	prev := head
	cur := head.Next
	var sum int
	for cur != nil {
		if cur.Val != 0 {
			sum += cur.Val
			cur = cur.Next
			continue
		}
		cur.Val = sum
		prev.Next = cur
		prev = cur
		cur = cur.Next
		sum = 0
	}
	return head.Next
}

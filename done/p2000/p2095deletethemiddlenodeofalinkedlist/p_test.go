package p2095deletethemiddlenodeofalinkedlist

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

func Test_deleteMiddle(t *testing.T) {
	for _, tc := range []struct {
		head string
		want string
	}{
		{"[1,3,4,7,1,2,6]", "[1,3,4,1,2,6]"},
		{"[1,2,3,4]", "[1,2,4]"},
		{"[2,1]", "[2]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			head := ParseList(tc.head)
			res := deleteMiddle(head)
			require.Equal(t, tc.want, res.String())
		})
	}
}

func deleteMiddle(head *ListNode) *ListNode {
	var n int
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	mid := n / 2
	cur = dummy
	for i := 0; i < mid; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

package p2046sortlinkedlistalreadysortedusingabsolutevalues

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

func Test_sortLinkedList(t *testing.T) {
	for _, tc := range []struct {
		list string
		want string
	}{
		{"[0,2,-5,5,10,-10]", "[-10,-5,0,2,5,10]"},
		{"[0,1,2]", "[0,1,2]"},
		{"[1]", "[1]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.list), func(t *testing.T) {
			list := ParseList(tc.list)
			res := sortLinkedList(list)
			require.Equal(t, tc.want, res.String())
		})
	}
}

func sortLinkedList(head *ListNode) *ListNode {

}

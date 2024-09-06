package p3217deletenodesfromlinkedelist

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_modifiedList(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		head string
		want string
	}{
		// {[]int{1, 2, 3}, "[1,2,3,4,5]", "[4,5]"},
		{[]int{1}, "[1,2,1,2,1,2]", "[2,2,2]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			x := ParseList(tc.head)
			s := modifiedList(tc.nums, x)
			require.Equal(t, tc.want, s.String())
		})
	}
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		m[x] = struct{}{}
	}
	root := &ListNode{
		Next: head,
	}
	cur := root
	for cur.Next != nil {
		if _, exists := m[cur.Next.Val]; exists {
			cur.Next = cur.Next.Next //snip
		} else {
			cur = cur.Next
		}
	}
	return root.Next
}

package p2130maximumtwinsumofalinkedlist

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

func Test_pairSum(t *testing.T) {
	for _, tc := range []struct {
		head string
		want int
	}{
		{"[5,4,2,1]", 6},
		{"[4,2,2,3]", 7},
		{"[1,100000]", 100001},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			h := ParseList(tc.head)
			require.Equal(t, tc.want, pairSum(h))
		})
	}
}

func pairSum(head *ListNode) int {
	// Check length of list
	var n int
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}

	// Move to n/2-1'th node, flipping the direction of the list
	cur, next := head, head.Next
	for i := 1; i < n/2; i++ {
		cur, next, next.Next = next, next.Next, cur
	}
	left, right := cur, next
	var maxSum int
	for i := 0; i < n/2; i++ {
		maxSum = max(maxSum, left.Val+right.Val)
		left = left.Next
		right = right.Next
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

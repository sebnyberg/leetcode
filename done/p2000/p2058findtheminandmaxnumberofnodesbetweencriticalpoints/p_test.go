package p2058findtheminandmaxnumberofnodesbetweencriticalpoints

import (
	"fmt"
	"log"
	"math"
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

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	for _, tc := range []struct {
		head string
		want []int
	}{
		{"[3,1]", []int{-1, -1}},
		{"[5,3,1,2,5,1,2]", []int{1, 3}},
		{"[1,3,2,2,3,2,2,2,7]", []int{3, 3}},
		{"[2,3,3,2]", []int{-1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			l := ParseList(tc.head)
			require.Equal(t, tc.want, nodesBetweenCriticalPoints(l))
		})
	}
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev, cur, next := head, head.Next, head.Next.Next
	firstCritical := -1
	prevCritical := -1
	maxDist := -1
	minDist := math.MaxInt32
	var i int
	for next != nil {
		if prev.Val < cur.Val && cur.Val > next.Val ||
			prev.Val > cur.Val && cur.Val < next.Val {
			// critical point
			if firstCritical != -1 {
				maxDist = i - firstCritical
				minDist = min(minDist, i-prevCritical)
			} else {
				firstCritical = i
			}
			prevCritical = i
		}
		i++
		prev, cur, next = cur, next, next.Next
	}
	if maxDist == -1 {
		return []int{-1, -1}
	}
	return []int{minDist, maxDist}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

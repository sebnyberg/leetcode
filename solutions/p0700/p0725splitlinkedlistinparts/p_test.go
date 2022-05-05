package p0725splitlinkedlistinparts

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

func Test_splitListToParts(t *testing.T) {
	for _, tc := range []struct {
		head string
		k    int
		want []string
	}{
		{"[1,2,3]", 5, []string{"[1]", "[2]", "[3]", "[]", "[]"}},
		{"[1,2,3,4,5,6,7,8,9,10]", 3, []string{"[1,2,3,4]", "[5,6,7]", "[8,9,10]"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			l := ParseList(tc.head)
			res := splitListToParts(l, tc.k)
			require.Equal(t, len(tc.want), len(res))
			for i, ll := range res {
				require.Equal(t, tc.want[i], ll.String())
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	// Count number of nodes
	cur := head
	var n int
	for cur != nil {
		n++
		cur = cur.Next
	}

	// Each group will take at least n/k elements, and the first n%k groups
	// will take one extra element.
	dummy := &ListNode{
		Next: head,
	}
	start := dummy.Next
	res := make([]*ListNode, 0, k)
	for i := 0; i < k; i++ {
		end := start
		groupCount := n / k
		if i < n%k {
			groupCount++
		}
		if groupCount == 0 {
			res = append(res, nil)
			continue
		}
		for j := 0; j < groupCount-1; j++ {
			end = end.Next
		}
		res = append(res, start)
		if end != nil {
			end.Next, start = nil, end.Next
		}
	}
	return res
}

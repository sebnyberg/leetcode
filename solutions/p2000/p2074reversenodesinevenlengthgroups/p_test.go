package p2074reversenodesinevenlengthgroups

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

func Test_reverseEvenLengthGroups(t *testing.T) {
	for _, tc := range []struct {
		listStr string
		want    string
	}{
		{"[0,4,2,1,3]", "[0,2,4,3,1]"},
		{"[5,2,6,3,9,1,7,3,8,4]", "[5,6,2,3,9,1,4,8,3,7]"},
		{"[1,1,0,6]", "[1,0,1,6]"},
		{"[2,1]", "[2,1]"},
		{"[8]", "[8]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.listStr), func(t *testing.T) {
			head := ParseList(tc.listStr)
			want := ParseList(tc.want)
			res := reverseEvenLengthGroups(head)
			require.Equal(t, want.String(), res.String())
		})
	}
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	listSlice := make([]*ListNode, 0, 100)
	cur := head
	for cur != nil {
		listSlice = append(listSlice, cur)
		cur = cur.Next
	}
	groupSize := 2
	var end int
	for start := 1; start < len(listSlice); start = end {
		end = min(start+groupSize, len(listSlice))
		groupSize++
		if (end-start)%2 != 0 {
			continue
		}
		for l, r := start, end-1; l < r; l, r = l+1, r-1 {
			listSlice[l], listSlice[r] = listSlice[r], listSlice[l]
		}
	}
	for i := 1; i < len(listSlice); i++ {
		listSlice[i-1].Next = listSlice[i]
	}
	listSlice[len(listSlice)-1].Next = nil
	return listSlice[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

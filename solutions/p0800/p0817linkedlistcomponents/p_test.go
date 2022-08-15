package p0817linkedlistcomponents

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

func Test_numComponents(t *testing.T) {
	for _, tc := range []struct {
		head string
		nums []int
		want int
	}{
		{"[0,1,2,3]", []int{0, 1, 3}, 2},
		{"[0,1,2,3,4]", []int{0, 3, 1, 4}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			head := ParseList(tc.head)
			res := numComponents(head, tc.nums)
			require.Equal(t, tc.want, res)
		})
	}
}

func numComponents(head *ListNode, nums []int) int {
	var exists [1e4 + 1]bool
	for _, x := range nums {
		exists[x] = true
	}
	var parent [1e4 + 1]int
	// var size [1e4 + 1]int
	for i := range parent {
		parent[i] = i
		// size[i] = 1
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		newParent := find(parent[a])
		parent[a] = newParent
		return parent[a]
	}
	union := func(a, b int) {
		if !exists[a] || !exists[b] {
			return
		}
		ra := find(a)
		rb := find(b)
		if ra != rb {
			// size[rb] += size[ra]
			parent[ra] = rb
		}
	}
	for curr := head; curr.Next != nil; curr = curr.Next {
		union(curr.Val, curr.Next.Val)
	}
	var seenRoot [1e4 + 1]bool
	var count int
	for _, x := range nums {
		root := find(x)
		if !seenRoot[root] {
			count++
		}
		seenRoot[root] = true
	}
	return count
}

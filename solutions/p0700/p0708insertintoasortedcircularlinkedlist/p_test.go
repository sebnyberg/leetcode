package p0708insertintoasortedcircularlinkedlist

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewLinkedListFromString(s string) *Node {
	if len(s) <= 2 {
		return nil
	}
	s = s[1 : len(s)-1]
	vals := strings.Split(s, ",")
	valInts := make([]int, len(vals))
	var err error
	for i := range vals {
		valInts[i], err = strconv.Atoi(vals[i])
		if err != nil {
			log.Fatalln(err)
		}
	}
	return NewLinkedList(valInts)
}

func NewLinkedList(items []int) *Node {
	head := &Node{Val: items[0]}
	cur := head
	for _, val := range items[1:] {
		cur.Next = &Node{Val: val}
		cur = cur.Next
	}
	cur.Next = head
	return head
}

func Test_Insert(t *testing.T) {
	for _, tc := range []struct {
		input     string
		insertVal int
		want      string
	}{
		{"[3,4,1]", 2, "[3,4,1,2]"},
		{"[]", 1, "[1]"},
		{"[1]", 0, "[1,0]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			input := NewLinkedListFromString(tc.input)
			res := insert(input, tc.insertVal)
			require.Equal(t, tc.want, res.String())
		})
	}
}

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) String() string {
	if n == nil {
		return "[]"
	}
	vals := make([]string, 1, 10)
	vals[0] = strconv.Itoa(n.Val)
	cur := n.Next
	for cur != n {
		vals = append(vals, strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	return "[" + strings.Join(vals, ",") + "]"
}

func insert(aNode *Node, x int) *Node {
	root := aNode // fix awful name
	if root == nil {
		res := &Node{Val: x}
		res.Next = res
		return res
	}

	// Go through list one round, trying to find a perfect insert position
	cur := root
	for {
		// Case 1: perfect insert position between nodes
		if cur.Val <= x && cur.Next.Val >= x {
			cur.Next = &Node{Val: x, Next: cur.Next}
			return root
		}
		// Case 2: About to reset to lower value
		if cur.Val > cur.Next.Val {
			if x >= cur.Val || x <= cur.Next.Val {
				cur.Next = &Node{Val: x, Next: cur.Next}
				return root
			}
		}
		cur = cur.Next
		if cur == root {
			break
		}
	}

	// Case 3: All values are the same, insert anywhere
	root.Next = &Node{Val: x, Next: root.Next}
	return root
}

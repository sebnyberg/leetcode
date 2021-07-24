package p0687longestunivaluepath

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTree(input string) *TreeNode {
	// Trim start/end []
	input = input[1 : len(input)-1]

	// Split by comma
	inputParts := strings.Split(input, ",")
	n := len(inputParts)

	if n == 0 || inputParts[0] == "" {
		return nil
	}

	// Create one node per element in the array
	nodes := make([]*TreeNode, n)
	for i, part := range inputParts {
		if part != "null" {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			nodes[i] = &TreeNode{Val: val}
		}
	}

	q := list.New()
	q.PushBack(nodes[0])

	i := 1
	for q.Len() > 0 && i < n {
		el := q.Remove(q.Front()).(*TreeNode)
		if nodes[i] != nil {
			el.Left = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
		if i >= n {
			break
		}
		if nodes[i] != nil {
			el.Right = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
	}

	return nodes[0]
}

func Test_longestUnivaluePath(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[5,4,5,1,1,5]", 2},
		{"[1,4,5,4,4,5]", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			res := longestUnivaluePath(root)
			require.Equal(t, tc.want, res)
		})
	}
}

func longestUnivaluePath(root *TreeNode) int {
	// For each node, return the longest path where that node is included, and
	// the longest sub-path with or without that node.
	if root == nil {
		return 0
	}
	var f pathFinder
	f.arrowLen(root)
	return f.maxLen
}

type pathFinder struct {
	maxLen int
}

func (f *pathFinder) arrowLen(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	leftLen := f.arrowLen(cur.Left)
	rightLen := f.arrowLen(cur.Right)
	totalLen := 0
	subLen := 0
	if cur.Left != nil && cur.Left.Val == cur.Val {
		leftLen++
		totalLen += leftLen
		subLen = max(subLen, leftLen)
	}
	if cur.Right != nil && cur.Right.Val == cur.Val {
		rightLen++
		totalLen += rightLen
		subLen = max(subLen, rightLen)
	}
	f.maxLen = max(f.maxLen, totalLen)
	return subLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

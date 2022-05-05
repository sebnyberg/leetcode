package p1026maximumdifferencebetweennodeandancestor

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

func Test_maxAncestorDiff(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[8,3,10,1,6,null,14,null,null,4,7,13]", 7},
		{"[1,null,2,null,0,3]", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, maxAncestorDiff(tree))
		})
	}
}

func maxAncestorDiff(root *TreeNode) int {
	// Simply traverse the tree, keeping track of the current max and min value
	// from ancestors to the current node. Then return the best result.
	return minMaxDiff(root, root.Val, root.Val)
}

func minMaxDiff(cur *TreeNode, minVal, maxVal int) int {
	if cur == nil {
		return -1
	}
	return max(
		minMaxDiff(cur.Left, min(minVal, cur.Val), max(maxVal, cur.Val)),
		max(
			minMaxDiff(cur.Right, min(minVal, cur.Val), max(maxVal, cur.Val)),
			max(abs(minVal-cur.Val), abs(maxVal-cur.Val)),
		),
	)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

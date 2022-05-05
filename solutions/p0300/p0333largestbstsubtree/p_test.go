package p0333largestbstsubtree

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_largestBSTSubtree(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[1,2]", 1},
		{"[10,5,15,1,8,null,7]", 3},
		{"[4,2,7,2,3,5,null,2,null,null,null,null,null,1]", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, largestBSTSubtree(tree))
		})
	}
}

func largestBSTSubtree(root *TreeNode) int {
	_, _, count := visit(root)
	return count
}

// Visit visits a node, returning the min and max imposed by its subtree, and
// the total number of nodes in the subtree.
func visit(cur *TreeNode) (int, int, int) {
	if cur == nil {
		// Why flip max/min? Because parents' check whether the min value is greater
		// than its max constraint and vice versa.
		return math.MaxInt32, math.MinInt32, 0
	}
	// While visiting this node, for a merge to happen, the max of the left
	// subtree must be smaller than the current value and min of the right
	// subtree. If this is the case, counts can be merged.
	leftMin, leftMax, leftCount := visit(cur.Left)
	rightMin, rightMax, rightCount := visit(cur.Right)
	if cur.Val > leftMax && cur.Val < rightMin {
		return min(cur.Val, leftMin), max(cur.Val, rightMax), leftCount + rightCount + 1
	}
	// This tree is not a valid subtree, so no parent trees can be either.
	// Create unreasonable constraints to invalidate all parent trees.
	return math.MinInt32, math.MaxInt32, max(leftCount, rightCount)
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

package p1373maximumsumbstinabinarytree

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

func Test_maxSumBST(t *testing.T) {
	for i, tc := range []struct {
		root string
		want int
	}{
		{"[4,3,null,1,2]", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, maxSumBST(tree))
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	// If, for each node, we know the range of values within its subtree, then
	// we can quickly validate whether the tree is a valid BST. The definition
	// states that the left range must end prior to the current node, and the
	// right range must start posterior to the current node.
	// If we collect the sum, the minimum value, and maximum value of each
	// subtree, then we can also determine the maximum sum in O(n).
	var res int
	visit(root, &res)
	return res
}

func visit(curr *TreeNode, res *int) (minVal, maxVal, sum int, validTree bool) {
	if curr == nil {
		return math.MaxInt32, math.MinInt32, 0, true
	}
	leftMin, leftMax, leftSum, leftValid := visit(curr.Left, res)
	rightMin, rightMax, rightSum, rightValid := visit(curr.Right, res)
	sum = leftSum + rightSum + curr.Val
	if leftMax < curr.Val && rightMin > curr.Val && leftValid && rightValid {
		validTree = true
		*res = max(*res, sum)
	}
	minVal = min(curr.Val, min(leftMin, rightMin))
	maxVal = max(curr.Val, max(leftMax, rightMax))
	return minVal, maxVal, sum, validTree
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

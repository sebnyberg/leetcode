package p0109convertsortarrtobst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBalanced(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
			&TreeNode{3, nil, nil},
		},
		&TreeNode{2, nil, nil},
	}
	// tree := &TreeNode{
	// 	1,
	// 	&TreeNode{
	// 		2,
	// 		&TreeNode{
	// 			3,
	// 			&TreeNode{4, nil, nil},
	// 			&TreeNode{4, nil, nil},
	// 		},
	// 		&TreeNode{3, nil, nil},
	// 	},
	// 	&TreeNode{3, nil, nil},
	// }

	require.Equal(t, false, isBalanced(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth, rightDepth := depth(root.Left, 1), depth(root.Right, 1)
	return abs(leftDepth-rightDepth) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(node *TreeNode, curDepth int) int {
	if node == nil {
		return curDepth - 1
	}
	return max(depth(node.Left, curDepth+1), depth(node.Right, curDepth+1))
}

package p0124bintreemaxpathsum

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxPathSum(t *testing.T) {
	// root := &TreeNode{
	// 	-10,
	// 	&TreeNode{9, nil, nil},
	// 	&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	// }

	// require.Equal(t, maxPathSum(root), 42)

	second := &TreeNode{
		-2,
		&TreeNode{1, nil, nil},
		nil,
	}
	require.Equal(t, maxPathSum(second), 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var f maxPathFinder
	f.maxPath = math.MinInt32
	open := f.findMaxSum(root)
	return max(f.maxPath, open)
}

type maxPathFinder struct {
	maxPath int
}

func (f *maxPathFinder) findMaxSum(node *TreeNode) int {
	f.maxPath = max(f.maxPath, node.Val)
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	var valLeft, valRight int
	if node.Left != nil {
		valLeft = f.findMaxSum(node.Left)
		f.maxPath = max(f.maxPath, node.Val+valLeft)
	}
	if node.Right != nil {
		valRight = f.findMaxSum(node.Right)
		f.maxPath = max(f.maxPath, node.Val+valRight)
	}
	if node.Left != nil && node.Right != nil {
		f.maxPath = max(f.maxPath, node.Val+valLeft+valRight)
	}
	return max(node.Val, max(node.Val+valLeft, node.Val+valRight))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

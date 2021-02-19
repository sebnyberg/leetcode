package p0124bintreemaxpathsum

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	_ = &TreeNode{
		-10,
		&TreeNode{9, nil, nil},
		&TreeNode{9, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}

	// require.Equal(t, maxPathSum(root), 42)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	return 0
}

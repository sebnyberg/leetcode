package p0101symmetrictree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestIsSymmetric(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			2,
			&TreeNode{4, nil, nil},
			&TreeNode{3, nil, nil},
		},
	}

	require.Equal(t, true, isSymmetric(tree))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return eq(root.Left, root.Right)
}

func eq(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		eq(left.Left, right.Right) &&
		eq(left.Right, right.Left)
}

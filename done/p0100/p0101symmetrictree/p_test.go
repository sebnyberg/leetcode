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
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	} else if right == nil {
		return left == nil
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Right, right.Left) && helper(left.Left, right.Right)
}

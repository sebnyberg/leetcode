package msft2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}
	res := lowestCommonAncestor(root, root.Left, root.Right)
	require.Equal(t, 6, res.Val)
	res = lowestCommonAncestor(root, root.Left, root.Left.Right)
	require.Equal(t, 2, res.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// The first node for which
	if p.Val > q.Val {
		p, q = q, p
	}
	// Both on right side
	if p.Val < root.Val {
		if q.Val >= root.Val {
			return root
		}
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val == root.Val {
		return root
	} else { // p.Val > root.Val
		if q.Val <= root.Val {
			return root
		}
		return lowestCommonAncestor(root.Right, p, q)
	}
}

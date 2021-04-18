package p0250countunivaluesubtrees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountUnivalSubtrees(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 5},
			Left:  &TreeNode{Val: 5},
		},
	}
	res := countUnivalSubtrees(root)
	require.Equal(t, 4, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	return 0
}

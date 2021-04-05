package p0543diameterofbtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	res := diameterOfBinaryTree(root)
	require.Equal(t, 3, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	p := &PathFinder{}
	p.longestPath(root)
	return p.maxLen
}

type PathFinder struct {
	maxLen int
}

func (f *PathFinder) longestPath(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	left := f.longestPath(cur.Left)
	right := f.longestPath(cur.Right)
	f.maxLen = max(f.maxLen, left+right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

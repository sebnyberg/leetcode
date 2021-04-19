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
	if root == nil {
		return 0
	}
	var res int
	if isUnival(root) {
		res++
	}
	l := countUnivalSubtrees(root.Left)
	r := countUnivalSubtrees(root.Right)
	return res + r + l
}

func isUnival(cur *TreeNode) bool {
	if cur == nil {
		return true
	}

	unival := true
	if cur.Left != nil {
		if cur.Left.Val == cur.Val {
			unival = isUnival(cur.Left)
		} else {
			unival = false
		}
	}
	if cur.Right != nil {
		if cur.Right.Val == cur.Val {
			unival = unival && isUnival(cur.Right)
		} else {
			unival = false
		}
	}
	return unival
}

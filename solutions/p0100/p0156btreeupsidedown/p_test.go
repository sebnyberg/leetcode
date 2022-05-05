package p0156btreeupsidedown

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpsideBT(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	_ = tree

	res := upsideDownBinaryTree(tree)
	require.Equal(t, 4, res.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	return flip(root, nil, nil)
}

func flip(cur *TreeNode, nextRight *TreeNode, nextLeft *TreeNode) *TreeNode {
	if cur == nil {
		return nextRight // previous root
	}
	curLeft, curRight := cur.Left, cur.Right
	cur.Left = nextLeft
	cur.Right = nextRight
	flip(curRight, nil, nil)
	return flip(curLeft, cur, curRight)
}

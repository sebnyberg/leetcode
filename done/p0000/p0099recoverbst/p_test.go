package p0099recoverbst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecoverTree(t *testing.T) {
	invalidTree := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil},
	}

	require.False(t, isValidBST(invalidTree))
	recoverTree(invalidTree)
	require.True(t, isValidBST(invalidTree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var h treeHealer
	h.search(root)
	h.first.Val, h.second.Val = h.second.Val, h.first.Val
}

type treeHealer struct {
	prev   *TreeNode
	first  *TreeNode
	second *TreeNode
}

func (t *treeHealer) search(cur *TreeNode) {
	if cur == nil {
		return
	}
	t.search(cur.Left)
	if t.prev != nil && t.prev.Val > cur.Val {
		if t.first == nil {
			t.first = t.prev
		}
		t.second = cur
	}
	t.prev = cur
	t.search(cur.Right)
}

// These were copied from 98 for testing purposes
func isValidBST(root *TreeNode) bool {
	return validate(root.Left, math.MinInt64, root.Val) && validate(root.Right, root.Val, math.MaxInt64)
}

func validate(node *TreeNode, minVal int, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Val >= maxVal || node.Val <= minVal {
		return false
	}
	return validate(node.Left, minVal, node.Val) && validate(node.Right, node.Val, maxVal)
}

package p0114flattenbsttolinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlatten(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			5,
			nil,
			&TreeNode{6, nil, nil},
		},
	}

	flatten(tree)
	require.Equal(t, true, true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	// Create a dummy prev-node to avoid the if-nil statement
	// If-statements (branches) reduce the efficiency of pipelined CPUs
	prev := &TreeNode{}

	var visit func(*TreeNode)
	visit = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		prev.Right = cur
		prev = cur
		// cur.Right will be overwritten by the next call to visit
		// store it as a function-local variable
		right := cur.Right
		visit(cur.Left)
		cur.Left = nil
		visit(right)
	}
	visit(root)
}

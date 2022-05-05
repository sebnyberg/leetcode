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
	if root == nil {
		return
	}
	var c nodeCollector
	c.collect(root)
	for i := 0; i < len(c.nodes)-1; i++ {
		c.nodes[i].Right = c.nodes[i+1]
		c.nodes[i].Left = nil
	}
	c.nodes[len(c.nodes)-1].Right = nil
	c.nodes[len(c.nodes)-1].Left = nil
}

type nodeCollector struct {
	nodes []*TreeNode
}

func (c *nodeCollector) collect(node *TreeNode) {
	if node == nil {
		return
	}
	c.nodes = append(c.nodes, node)
	c.collect(node.Left)
	c.collect(node.Right)
}

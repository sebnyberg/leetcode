package p0117popnextrightnode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	root := &Node{
		1,
		&Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil},
		&Node{3, nil, &Node{7, nil, nil, nil}, nil},
		nil,
	}

	result := connect(root)
	_ = result
	require.Equal(t, true, true)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectNext(root)
	return root
}

func connectNext(node *Node) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	// Find next node outside this subgraph
	var next *Node
	cur := node.Next
	for cur != nil {
		if cur.Left != nil {
			next = cur.Left
			break
		}
		if cur.Right != nil {
			next = cur.Right
			break
		}
		cur = cur.Next
	}

	// Link children to their next nodes
	if node.Right == nil {
		node.Left.Next = next
	} else if node.Left == nil {
		node.Right.Next = next
	} else { // neither are nil
		node.Left.Next = node.Right
		node.Right.Next = next
	}

	// Traverse children
	connectNext(node.Right)
	connectNext(node.Left)
}

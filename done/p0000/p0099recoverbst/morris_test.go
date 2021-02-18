package p0099recoverbst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecoverTreeMorris(t *testing.T) {
	// invalidTree := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 4,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 	},
	// }

	invalidTree := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil},
	}

	require.False(t, isValidBST(invalidTree))
	recoverTreeMorris(invalidTree)
	require.True(t, isValidBST(invalidTree))
}

func recoverTreeMorris(root *TreeNode) {
	current := root
	var preInorder *TreeNode
	var first, second *TreeNode

	// Morris in-order traversal
	for current != nil {
		if current.Left == nil {
			// Follow the link to the right,
			// Either originally in the graph, or created by the threading below
			if preInorder != nil && preInorder.Val > current.Val {
				// current is the offending node.
				// pre might be the offending node too.
				if first == nil {
					first = preInorder
				}
				second = current
			}
			preInorder = current
			current = current.Right
			continue
		}
		// Find the inorder predecessor of current
		curPredecessor := current.Left
		for curPredecessor.Right != nil && curPredecessor.Right != current {
			curPredecessor = curPredecessor.Right
		}

		// pre.Right should always point to its successor, so if it is nil,
		// the node should be threaded to the current
		if curPredecessor.Right == nil {
			curPredecessor.Right = current
			current = current.Left
		} else {
			curPredecessor.Right = nil
			if preInorder != nil && preInorder.Val > current.Val {
				// current is the offending node.
				// pre might be the offending node too.
				if first == nil {
					first = preInorder
				}
				second = current
			}
			preInorder = current
			current = current.Right
		}
	}

	// Do swap with first, second
	first.Val, second.Val = second.Val, first.Val
}

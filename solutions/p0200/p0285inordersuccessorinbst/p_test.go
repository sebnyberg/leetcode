package p0285inordersuccessorinbst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInorderSuccessor(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6},
	}
	p := &TreeNode{Val: 6}

	res := inorderSuccessor(root, p)
	require.Equal(t, nil, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	switch {
	case root == nil, p == nil:
		return nil
	case root.Val <= p.Val:
		return inorderSuccessor(root.Right, p)
	default: // root.Val > p.Val
		minLeft := inorderSuccessor(root.Left, p)
		if minLeft != nil && minLeft.Val < root.Val {
			return minLeft
		}
		return root
	}
}

package p0129sumroottoleaf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		4,
		&TreeNode{
			9,
			&TreeNode{5, nil, nil},
			&TreeNode{1, nil, nil},
		},
		&TreeNode{0, nil, nil},
	}

	require.Equal(t, 1026, sumNumbers(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var s pathSummer
	if root == nil {
		return 0
	}
	s.sumPath(root, 0)
	return s.sum
}

type pathSummer struct {
	sum int
}

func (s *pathSummer) sumPath(node *TreeNode, curSum int) {
	if node == nil {
		return
	}
	curSum = curSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		s.sum += curSum
		return
	}
	if node.Left != nil {
		s.sumPath(node.Left, curSum)
	}
	if node.Right != nil {
		s.sumPath(node.Right, curSum)
	}
}

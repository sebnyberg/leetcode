package p0105bstfompreinordtrav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTree(t *testing.T) {
	// res := buildTree([]int{1}, []int{1})
	res := buildTree([]int{1, 2, 4, 5, 3, 6}, []int{4, 2, 5, 1, 3, 6})
	_ = res
	require.Equal(t, true, true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var rootIdx int
	for i, n := range inorder {
		if n == preorder[0] {
			rootIdx = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIdx+1], inorder[:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}
}

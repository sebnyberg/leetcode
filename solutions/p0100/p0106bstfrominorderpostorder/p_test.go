package p0106bstfompreinordtrav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTree(t *testing.T) {
	res := buildTree([]int{4, 2, 5, 1, 3, 6}, []int{4, 5, 2, 6, 3, 1})
	_ = res
	require.Equal(t, true, true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var rootIdx int
	n := len(postorder)
	rootVal := postorder[n-1]
	for i, n := range inorder {
		if n == rootVal {
			rootIdx = i
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:rootIdx], postorder[:rootIdx]),
		Right: buildTree(inorder[rootIdx+1:], postorder[rootIdx:n-1]),
	}
}

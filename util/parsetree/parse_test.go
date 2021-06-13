package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTree(t *testing.T) {
	// in := "[0,1,4,2,8,37,13,3,34,14,29,66,45,22,19,5,6,39,69,null,17,null,35,null,null,null,null,32,null,null,58,null,9,10,7,null,55,89,null,42,51,57,null,86,null,null,null,11,18,53,15,12,null,null,null,null,null,48,null,80,84,75,65,null,null,26,64,27,21,9]"
	in := "[5,4,7,3,null,2,null,-1,null,9]"
	// in := "[0,1,4]"
	// in := "[0,1,4,2,8,37,13,3,34,14,29,66,null,null,null]"
	root := ParseTree(in)
	// inorder := []int{26, 11, 64, 9, 27, 18, 21, 5, 3, 9, 53, 10, 15, 6, 12, 7, 2, 55, 39, 34, 89, 69, 1, 48, 42, 17, 80, 51, 84, 14, 8, 75, 57, 65, 35, 29}
	inorder := []int{-1, 3, 4, 5, 9, 2, 7}
	// inorder := []int{1, 0, 4}
	// inorder := []int{3, 2, 34, 1, 14, 8, 29, 0, 66, 37, 4, 13}
	res := make([]int, 0, len(inorder)*2)
	collectOrder(root, &res)
	require.Equal(t, inorder, res)
	_ = root
}

func collectOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	collectOrder(root.Left, res)
	*res = append(*res, root.Val)
	collectOrder(root.Right, res)
}

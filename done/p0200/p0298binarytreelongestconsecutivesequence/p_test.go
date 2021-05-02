package p0298longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
	}
	res := longestConsecutive(root)
	_ = res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 1)
}

func helper(root *TreeNode, len int) int {
	left, right := len, len
	if root.Left != nil {
		if root.Left.Val == root.Val+1 {
			left = helper(root.Left, len+1)
		} else {
			left = helper(root.Left, 1)
		}
	}
	if root.Right != nil {
		if root.Right.Val == root.Val+1 {
			right = helper(root.Right, len+1)
		} else {
			right = helper(root.Right, 1)
		}
	}
	return max(len, max(left, right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

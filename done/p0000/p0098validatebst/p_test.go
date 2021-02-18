package p0098validatebst

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root.Left, math.MinInt64, root.Val) && validate(root.Right, root.Val, math.MaxInt64)
}

func validate(node *TreeNode, minVal int, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Val >= maxVal || node.Val <= minVal {
		return false
	}
	return validate(node.Left, minVal, node.Val) && validate(node.Right, node.Val, maxVal)
}

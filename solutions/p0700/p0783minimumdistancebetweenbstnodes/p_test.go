package p0783minimumdistancebetweenbstnodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res := visit(root, math.MinInt32, math.MaxInt32)
	return res
}

func visit(node *TreeNode, minVal, maxVal int) int {
	if node == nil {
		return math.MaxInt32
	}
	res := min(node.Val-minVal, maxVal-node.Val)
	res = min(res, visit(node.Left, minVal, node.Val))
	res = min(res, visit(node.Right, node.Val, maxVal))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

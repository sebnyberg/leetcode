package p1080insufficientnodesinroottoleafpaths

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	res, _ := f(root, 0, limit)
	return res
}

func f(curr *TreeNode, sum, limit int) (*TreeNode, int) {
	sum += curr.Val
	maxSum := math.MinInt32
	if curr.Left == nil && curr.Right == nil {
		maxSum = sum
	}
	if curr.Left != nil {
		left, leftSum := f(curr.Left, sum, limit)
		curr.Left = left
		maxSum = max(maxSum, leftSum)
	}
	if curr.Right != nil {
		right, rightSum := f(curr.Right, sum, limit)
		curr.Right = right
		maxSum = max(maxSum, rightSum)
	}
	if maxSum >= limit {
		return curr, maxSum
	}

	return nil, maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

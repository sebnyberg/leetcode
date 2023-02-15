package p1161maximumlevelsumofabinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	// Just do sums per level with BFS and find the maximum sum.
	//
	curr := []*TreeNode{root}
	next := []*TreeNode{}

	maxSum := math.MinInt32
	var res int

	for level := 1; len(curr) > 0; level++ {
		var sum int
		next = next[:0]
		for _, x := range curr {
			if x.Left != nil {
				next = append(next, x.Left)
			}
			if x.Right != nil {
				next = append(next, x.Right)
			}
			sum += x.Val
		}
		if sum > maxSum {
			maxSum = sum
			res = level
		}
		curr, next = next, curr
	}
	return res
}

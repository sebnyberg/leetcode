package p2583kthlargestsuminabinarytree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	// Just collect sums, sort, and return
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	var sums []int
	for len(curr) > 0 {
		next = next[:0]
		var sum int
		for _, x := range curr {
			sum += x.Val
			if x.Left != nil {
				next = append(next, x.Left)
			}
			if x.Right != nil {
				next = append(next, x.Right)
			}
		}
		sums = append(sums, sum)
		curr, next = next, curr
	}
	if k > len(sums) {
		return -1
	}
	sort.Ints(sums)
	return int64(sums[len(sums)-k])
}

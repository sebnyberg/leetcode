package p2476closestnodesqueriesinabinarysearchtree

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	// Because there are so many annoying edge-cases, I collect all nodes in a
	// sorted list (with sentinel values at the first and last element).
	// Then, match the values found in the tree against queries sorted by value.
	vals := []int{math.MinInt32}
	var collect func(cur *TreeNode)
	collect = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		collect(cur.Left)
		vals = append(vals, cur.Val)
		collect(cur.Right)
	}
	collect(root)
	vals = append(vals, math.MaxInt32)

	n := len(queries)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return queries[idx[i]] < queries[idx[j]]
	})
	res := make([][]int, n)
	var j int
	for i := range queries {
		for queries[idx[i]] > vals[j] {
			j++
		}
		if queries[idx[i]] == vals[j] {
			res[idx[i]] = []int{vals[j], vals[j]}
			continue
		}
		res[idx[i]] = []int{vals[j-1], vals[j]}
	}
	for i := range res {
		for j := range res[i] {
			if res[i][j] == math.MinInt32 {
				res[i][j] = -1
			}
			if res[i][j] == math.MaxInt32 {
				res[i][j] = -1
			}
		}
	}
	return res
}

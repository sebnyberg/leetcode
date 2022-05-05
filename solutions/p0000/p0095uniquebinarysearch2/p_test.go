package p0095uniquebinarysearch2

import (
	"testing"
)

func Test_generateTrees(t *testing.T) {
	// res := generateTrees(3)
	// _ = res
	// require.Equal(t, 1, 2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	// For each number in n, e.g. m
	// Given that m belongs to the root, create one tree for
	// each combination of the subtrees formed by the numbers
	// remaining on each side.
	// For example, given n=5, and m=3, create one tree for
	// each tree formed in the interval [1,2] (left side)
	// and [4,5] (right side), which adds up to 4 trees in total with m=3 at the root
	return findSubTrees(1, n)
}

func findSubTrees(from, to int) []*TreeNode {
	if from > to {
		return []*TreeNode{nil}
	}
	if from == to {
		return []*TreeNode{
			{
				Val:   from,
				Left:  nil,
				Right: nil,
			},
		}
	}

	res := make([]*TreeNode, 0)
	for i := from; i <= to; i++ {
		lefts := findSubTrees(from, i-1)
		rights := findSubTrees(i+1, to)
		for _, left := range lefts {
			for _, right := range rights {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return res
}

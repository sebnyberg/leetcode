package fb1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger interface {
	// Return true if this NestedInteger holds a single integer, rather than a nested list.
	IsInteger() bool

	// Return the single integer that this NestedInteger holds, if it holds a single integer
	// The result is undefined if this NestedInteger holds a nested list
	// So before calling this method, you should have a check
	GetInteger() int

	// Set this NestedInteger to hold a single integer.
	SetInteger(value int)

	// Set this NestedInteger to hold a nested list and adds a nested integer to it.
	Add(elem NestedInteger)

	// Return the nested list that this NestedInteger holds, if it holds a nested list
	// The list length is zero if this NestedInteger holds a single integer
	// You can access NestedInteger's List element directly if you want to modify it
	GetList() []NestedInteger
}

// type IntAndDepth struct {
// 	val   int
// 	depth int
// }

// func depthSumInverse(nestedList []*NestedInteger) int {
// 	// Traverse list of nested integers
// 	f := ListFlattener{
// 		values: make([]IntAndDepth, 1, len(nestedList)),
// 	}
// 	f.traverse(0, nestedList)
// 	var res int
// 	for _, el := range f.values {
// 		res += el.val * (f.maxDepth - el.depth + 1)
// 	}
// 	return res
// }

// type ListFlattener struct {
// 	values   []IntAndDepth
// 	maxDepth int
// }

// func (f *ListFlattener) traverse(curDepth int, nestedList []*NestedInteger) {
// 	f.maxDepth = max(f.maxDepth, curDepth)
// 	for _, l := range nestedList {
// 		if l.IsInteger() {
// 			f.values = append(f.values, IntAndDepth{l.GetInteger(), curDepth})
// 		} else {
// 			// Traverse further down
// 			f.traverse(curDepth+1, l.GetList())
// 		}
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func Test_rangeSumBST(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Right: &TreeNode{Val: 18},
		},
	}

	res := rangeSumBST(tree, 7, 15)
	require.Equal(t, 32, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var curVal int
	if root.Val <= high && root.Val >= low {
		curVal += root.Val
	}
	if low < root.Val {
		// go left
		curVal += rangeSumBST(root.Left, low, high)
	}
	if high > root.Val {
		// go right
		curVal += rangeSumBST(root.Right, low, high)
	}
	return curVal
}

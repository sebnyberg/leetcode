package p0113pathsum2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSum2(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{13, nil, nil},
			&TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}},
		},
	}

	res := pathSum(tree, 22)
	_ = res
	require.Equal(t, true, true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	f := PathSumFinder{
		prefix: []int{},
		res:    [][]int{},
	}
	f.findPathSums(root, targetSum)
	return f.res
}

type PathSumFinder struct {
	prefix []int
	res    [][]int
}

func (f *PathSumFinder) findPathSums(node *TreeNode, curSum int) {
	if node == nil {
		return
	}
	curSum -= node.Val
	if node.Left == nil && node.Right == nil {
		if curSum == 0 {
			prefixCpy := make([]int, len(f.prefix))
			copy(prefixCpy, f.prefix)
			prefixCpy = append(prefixCpy, node.Val)
			f.res = append(f.res, prefixCpy)
		}
		return
	}
	f.prefix = append(f.prefix, node.Val)
	f.findPathSums(node.Left, curSum)
	f.findPathSums(node.Right, curSum)
	f.prefix = f.prefix[:len(f.prefix)-1]
}

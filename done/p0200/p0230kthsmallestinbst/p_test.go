package p0230kthsmallestinbst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	require.Equal(t, 1, kthSmallest(root, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	f := indexFinder{0, k, 0}
	f.find(root)
	return f.res
}

type indexFinder struct {
	i   int
	k   int
	res int
}

func (f *indexFinder) find(cur *TreeNode) {
	if f.i > f.k {
		return
	}
	if cur.Left != nil {
		f.find(cur.Left)
	}
	f.i++
	if f.i == f.k {
		f.res = cur.Val
		return
	}
	if cur.Right != nil {
		f.find(cur.Right)
	}
}

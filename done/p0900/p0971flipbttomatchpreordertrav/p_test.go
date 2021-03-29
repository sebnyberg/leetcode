package p0971flipbttomatchpreordertrav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_flipMatchVoyage(t *testing.T) {
	for _, tc := range []struct {
		name   string
		root   *TreeNode
		voyage []int
		want   []int
	}{
		{
			"Valid flip (one needed)",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			[]int{1, 3, 2},
			[]int{1},
		},
		{
			"No flips needed",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			[]int{1, 2, 3},
			[]int{},
		},
		{
			"Invalid flip",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			[]int{2, 1},
			[]int{-1},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, flipMatchVoyage(tc.root, tc.voyage))
		})
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	f := treeFlipper{
		flippedNodes: make([]int, 0),
	}
	f.visit(root, voyage)
	if f.idx == len(voyage) {
		return f.flippedNodes
	}
	return []int{-1}
}

type treeFlipper struct {
	flippedNodes []int
	idx          int
}

func (f *treeFlipper) visit(cur *TreeNode, voyage []int) {
	if cur.Val != voyage[f.idx] {
		return
	}
	f.idx++
	if f.idx > len(voyage) {
		return
	}
	if cur.Left != nil && cur.Left.Val != voyage[f.idx] {
		// Flip (does not ensure right order, that is checked above)
		cur.Left, cur.Right = cur.Right, cur.Left
		f.flippedNodes = append(f.flippedNodes, cur.Val)
	}
	if cur.Left != nil {
		f.visit(cur.Left, voyage)
	}
	if cur.Right != nil {
		f.visit(cur.Right, voyage)
	}
}

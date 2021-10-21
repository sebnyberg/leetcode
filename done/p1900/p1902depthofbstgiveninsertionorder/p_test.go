package p1902depthofbstgiveninsertionorder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepthBST(t *testing.T) {
	for _, tc := range []struct {
		order []int
		want  int
	}{
		{[]int{2, 1, 4, 3}, 3},
		{[]int{2, 1, 3, 4}, 3},
		{[]int{1, 2, 3, 4}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.order), func(t *testing.T) {
			require.Equal(t, tc.want, maxDepthBST(tc.order))
		})
	}
}

type treeNode struct {
	left, right *treeNode
	val         int
}

func (n *treeNode) insert(x, depth int) int {
	if x < n.val {
		if n.left != nil {
			return n.left.insert(x, depth+1)
		} else {
			n.left = &treeNode{val: x}
			return depth + 1
		}
	} else {
		if n.right != nil {
			return n.right.insert(x, depth+1)
		} else {
			n.right = &treeNode{val: x}
			return depth + 1
		}
	}
}

func maxDepthBST(order []int) int {
	n := len(order)

	// Insert steadily increasing / decreasing numbers into the BST before
	// doing a regular insert.
	root := &treeNode{val: order[0]}

	seen := make([]bool, n)
	seen[0] = true

	right := root
	prevRight := order[0]
	depthRight := 1
	left := root
	prevLeft := order[0]
	depthLeft := 1
	for i, num := range order {
		if num > prevRight {
			prevRight = num
			right.right = &treeNode{val: num}
			right = right.right
			seen[i] = true
			depthRight++
		} else if num < prevLeft {
			prevLeft = num
			left.left = &treeNode{val: num}
			left = left.left
			seen[i] = true
			depthLeft++
		}
	}

	maxDepth := max(depthRight, depthLeft)

	// Finally, do a regular BST insert of remaining values
	for i, num := range order {
		if seen[i] {
			continue
		}
		maxDepth = max(maxDepth, root.insert(num, 1))
	}

	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1522

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiameter(t *testing.T) {
	tree := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2},
			{
				Val: 3,
				Children: []*Node{
					{Val: 6},
					{
						Val: 7,
						Children: []*Node{
							{
								Val:      11,
								Children: []*Node{{Val: 14}},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{
							{
								Val: 12,
							},
						}},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
						Children: []*Node{
							{
								Val: 13,
							},
						},
					},
					{
						Val: 10,
					},
				},
			},
		},
	}
	res := diameter(tree)
	require.Equal(t, 7, res)
}

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	// The longest diameter is the longest pair of legs found in the tree
	res, _ := visit(root)
	return res
}

// visit returns the maximum diameter for the provided node, and the longest
// depth for any path where the node is the root.
func visit(node *Node) (int, int) {
	if node == nil {
		return 0, 0
	}
	var maxDiam, maxDepth, secondMaxDepth int
	for _, child := range node.Children {
		diam, depth := visit(child)
		if diam > maxDiam {
			maxDiam = diam
		}
		d := depth + 1
		if d > maxDepth {
			maxDepth, secondMaxDepth = d, maxDepth
		} else if d >= secondMaxDepth {
			secondMaxDepth = d
		}
	}
	return max(maxDiam, maxDepth+secondMaxDepth), maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

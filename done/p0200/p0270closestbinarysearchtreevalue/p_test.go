package p0270closestbinarysearchtreevalue

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClosestValue(t *testing.T) {
	// root := &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val:  2,
	// 		Left: &TreeNode{Val: 1},
	// 	},
	// 	Right: &TreeNode{Val: 5},
	// }
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	res := closestValue(root, 3.714286)
	require.Equal(t, 2, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	minDelta := math.MaxFloat32
	cur := root
	res := root.Val
	for cur != nil {
		v := float64(cur.Val)
		d := abs(v - target)
		if d < minDelta {
			minDelta = d
			res = cur.Val
		}
		if v < target {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	return res
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

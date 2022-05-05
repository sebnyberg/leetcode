package p0968binarytreecamera

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCameraCover(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		root := &TreeNode{
			Left: &TreeNode{
				Right: &TreeNode{
					Left: &TreeNode{
						Right: &TreeNode{
							Left: &TreeNode{},
						},
					},
				},
			},
		}
		res := minCameraCover(root)
		require.Equal(t, 2, res)
	})

	t.Run("second", func(t *testing.T) {
		root := &TreeNode{
			Left: &TreeNode{
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
		}
		res := minCameraCover(root)
		require.Equal(t, 1, res)
	})

	t.Run("third", func(t *testing.T) {
		root := &TreeNode{}
		res := minCameraCover(root)
		require.Equal(t, 1, res)
	})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	covered := make(map[*TreeNode]bool)
	return dfs(root, nil, covered)
}

func dfs(cur *TreeNode, par *TreeNode, covered map[*TreeNode]bool) int {
	if cur == nil {
		return 0
	}
	left := dfs(cur.Left, cur, covered)
	right := dfs(cur.Right, cur, covered)
	if par == nil && !covered[cur] {
		covered[cur] = true
		return left + right + 1
	}
	if (cur.Left != nil && !covered[cur.Left]) || (cur.Right != nil && !covered[cur.Right]) {
		covered[cur.Left] = true
		covered[cur.Right] = true
		covered[cur] = true
		covered[par] = true
		return left + right + 1
	}
	return left + right
}

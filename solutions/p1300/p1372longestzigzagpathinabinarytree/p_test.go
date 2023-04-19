package p1372longestzigzagpathinabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	// A new path starts as a new direction is taken.
	// We may simply go down all paths and see which one is deepest.
	var res int
	longestZigZagDir(root, true, 0, &res)
	return res
}

func longestZigZagDir(cur *TreeNode, wantLeft bool, depth int, res *int) {
	if cur == nil {
		return
	}
	*res = max(*res, depth)
	if wantLeft {
		longestZigZagDir(cur.Left, false, depth+1, res)
		longestZigZagDir(cur.Right, true, 1, res)
	} else {
		longestZigZagDir(cur.Right, true, depth+1, res)
		longestZigZagDir(cur.Left, false, 1, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0687longestunivaluepath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	visit(root, &res)
	return res - 1
}

// visit returns the longest univariate path starting in this node and checks
// for longest total path with the root in this node.
func visit(cur *TreeNode, res *int) int {
	if cur == nil {
		return 0
	}
	left := visit(cur.Left, res)
	right := visit(cur.Right, res)
	tot := 1
	single := 1
	if cur.Left != nil && cur.Left.Val == cur.Val {
		single = max(single, 1+left)
		tot += left
	}
	if cur.Right != nil && cur.Right.Val == cur.Val {
		single = max(single, 1+right)
		tot += right
	}
	*res = max(*res, tot)
	return single
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2331evaluatebooleanbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	return eval(root)
}

func eval(cur *TreeNode) bool {
	if cur.Val <= 1 {
		return cur.Val == 1
	}
	left := eval(cur.Left)
	right := eval(cur.Right)
	if cur.Val == 2 {
		return left || right
	}
	return left && right
}

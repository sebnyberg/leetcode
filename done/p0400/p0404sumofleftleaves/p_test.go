package p0404sumofleftleaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root, false)
}
func helper(cur *TreeNode, left bool) int {
	if cur == nil {
		return 0
	}
	if cur.Left == nil && cur.Right == nil && left {
		return cur.Val
	}
	return helper(cur.Left, true) + helper(cur.Right, false)
}

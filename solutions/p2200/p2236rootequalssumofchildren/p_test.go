package p2236rootequalssumofchildren

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == sum(root.Left)+sum(root.Right)
}

func sum(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return sum(cur.Left) + sum(cur.Right) + cur.Val
}

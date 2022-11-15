package p0965univaluedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return ok(root, root.Val)
}
func ok(curr *TreeNode, val int) bool {
	if curr == nil {
		return true
	}
	return curr.Val == val && ok(curr.Left, val) && ok(curr.Right, val)
}

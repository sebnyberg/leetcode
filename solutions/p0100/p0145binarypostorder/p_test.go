package p0145binarypreorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(
		append(postorderTraversal(root.Left), postorderTraversal(root.Right)...),
		root.Val,
	)
}

package p0144binarypreorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(
		append([]int{root.Val}, preorderTraversal(root.Left)...),
		preorderTraversal(root.Right)...,
	)
}

package p0112pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

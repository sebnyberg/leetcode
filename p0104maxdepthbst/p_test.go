package p0104maxdepthbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	return findMaxDepth(root, 1)
}

func findMaxDepth(node *TreeNode, level int) int {
	if node == nil {
		return level - 1
	}
	return max(findMaxDepth(node.Left, level+1), findMaxDepth(node.Right, level+1))
}

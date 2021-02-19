package p0111mindepthbintree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return minDepthLevel(root, 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepthLevel(node *TreeNode, level int) int {
	if node == nil {
		return level - 1
	}
	if node.Left == nil {
		return minDepthLevel(node.Right, level+1)
	} else if node.Right == nil {
		return minDepthLevel(node.Left, level+1)
	} else {
		return min(minDepthLevel(node.Left, level+1), minDepthLevel(node.Right, level+1))
	}
}

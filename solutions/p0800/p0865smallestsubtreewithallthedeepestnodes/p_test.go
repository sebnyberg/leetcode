package p0865smallestsubtreewithallthedeepestnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := visit(root, 0)
	return res
}

func visit(cur *TreeNode, level int) (int, *TreeNode) {
	if cur == nil {
		return -1, nil
	}
	left, leftNode := visit(cur.Left, level+1)
	right, rightNode := visit(cur.Right, level+1)
	sides := max(left, right)
	if level > sides {
		return level, cur
	}
	if left == right {
		return sides, cur
	}
	if left > right {
		return left, leftNode
	}
	return right, rightNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

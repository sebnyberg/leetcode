package p1145binarytreecoloringgame

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	// First, count all nodes so that we can calculate parent-edge nodes based
	// on subtree counts.
	var subTreeCount [2]int
	rootCount := countNodes(root, x, &subTreeCount)
	left := subTreeCount[0]
	right := subTreeCount[1]
	parent := rootCount - 1 - left - right
	return left > parent+right+1 ||
		right > parent+left+1 ||
		parent > left+right+1
}

func countNodes(curr *TreeNode, x int, res *[2]int) int {
	if curr == nil {
		return 0
	}
	left := countNodes(curr.Left, x, res)
	right := countNodes(curr.Right, x, res)
	if curr.Val == x {
		*res = [2]int{left, right}
	}
	return 1 + left + right
}

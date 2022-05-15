package p2265countnodesequaltoaverageofsubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var res int
	visit(root, &res)
	return res
}

func visit(cur *TreeNode, res *int) (sum, count int) {
	if cur == nil {
		return 0, 0
	}
	ls, lc := visit(cur.Left, res)
	rs, rc := visit(cur.Right, res)
	subSum := ls + rs + cur.Val
	subCount := lc + rc + 1
	avg := subSum / subCount
	if cur.Val == avg {
		*res++
	}
	return subSum, subCount
}

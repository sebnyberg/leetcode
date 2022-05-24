package p2265countnodesequaltoaverageofsubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var res int
	countSubtree(root, &res)
	return res
}

func countSubtree(cur *TreeNode, res *int) (count int, sum int) {
	if cur == nil {
		return 0, 0
	}
	leftCount, leftSum := countSubtree(cur.Left, res)
	rightCount, rightSum := countSubtree(cur.Right, res)
	totalCount := leftCount + rightCount + 1
	totalSum := leftSum + rightSum + cur.Val
	if (totalSum / totalCount) == cur.Val {
		*res++
	}

	return totalCount, totalSum
}

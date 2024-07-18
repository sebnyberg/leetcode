package p1530numberofgoodleafnodespairs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	// At each node, gather a binning of the distances to all leaves in the
	// subtree. Then for each pair of bins, check if the sum of the bins is
	// less than or equal to the distance.
	var res int
	accumulateBins(root, distance, &res)
	return res
}

func accumulateBins(root *TreeNode, distance int, res *int) [10]int {
	if root == nil {
		return [10]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [10]int{1: 1}
	}
	left := accumulateBins(root.Left, distance, res)
	right := accumulateBins(root.Right, distance, res)

	for i := 1; i <= 9; i++ {
		for j := 1; i+j <= distance; j++ {
			*res += left[i] * right[j]
		}
	}

	var next [10]int
	for i := 9; i >= 1; i-- {
		next[i] = left[i-1] + right[i-1]
	}

	return next
}

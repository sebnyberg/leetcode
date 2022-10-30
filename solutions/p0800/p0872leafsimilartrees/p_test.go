package p0872leafsimilartrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves func(curr *TreeNode, res *[]int)
	leaves = func(curr *TreeNode, res *[]int) {
		if curr == nil {
			return
		}
		leaves(curr.Left, res)
		if curr.Left == nil && curr.Right == nil {
			*res = append(*res, curr.Val)
		}
		leaves(curr.Right, res)
	}
	var first []int
	var second []int
	leaves(root1, &first)
	leaves(root2, &second)
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

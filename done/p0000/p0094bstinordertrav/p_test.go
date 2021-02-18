package p0094bstinordertrav

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traverse(root, &res)
	return res
}

func traverse(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	traverse(cur.Left, res)
	*res = append(*res, cur.Val)
	traverse(cur.Right, res)
}

package p0199binarytreeright

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	traverse(root, 0, &res)
	return res
}

func traverse(node *TreeNode, level int, res *[]int) {
	if node == nil {
		return
	}
	if len(*res) < level+1 {
		*res = append(*res, node.Val)
	}
	traverse(node.Right, level+1, res)
	traverse(node.Left, level+1, res)
}

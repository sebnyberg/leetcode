package p0623addonerowtotree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		root = &TreeNode{v, root, nil}
		return root
	}
	if root == nil {
		return nil
	}
	buildTree(root, v, d)
	return root
}

func buildTree(node *TreeNode, v int, d int) {
	if node == nil {
		return
	}
	if d == 2 && node != nil {
		node.Left = &TreeNode{v, node.Left, nil}
		node.Right = &TreeNode{v, nil, node.Right}
	}
	addOneRow(node.Left, v, d-1)
	addOneRow(node.Right, v, d-1)
}

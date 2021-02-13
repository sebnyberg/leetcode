package p1038bsttogst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	v := 0
	c(root, &v)
	return root
}

func c(node *TreeNode, val *int) {
	// Visit right node
	if node.Right != nil {
		c(node.Right, val)
	}
	// Assign value to current node
	*val += node.Val
	node.Val = *val
	// Visit left nodes
	if node.Left != nil {
		c(node.Left, val)
	}
}

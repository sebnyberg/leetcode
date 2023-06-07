package p1382balanceabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	collect(root, &nodes)
	return balance(nodes)
}

func balance(nodes []*TreeNode) *TreeNode {
	n := len(nodes)
	if n == 0 {
		return nil
	}
	mid := n / 2
	res := &TreeNode{
		Val:   nodes[mid].Val,
		Left:  balance(nodes[:mid]),
		Right: balance(nodes[mid+1:]),
	}
	return res
}

func collect(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	collect(root.Left, nodes)
	*nodes = append(*nodes, root)
	collect(root.Right, nodes)
}

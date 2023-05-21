package p1315sumofnodeswithevenvaluedgrandparent

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	res := visit(root, 0)
	return res
}

func visit(cur *TreeNode, parent int) int {
	if cur == nil {
		return 0
	}
	var res int
	if parent&(1<<2) > 0 {
		res += cur.Val
	}
	if cur.Val%2 == 0 {
		parent |= 1
	}
	return res + visit(cur.Left, parent<<1) + visit(cur.Right, parent<<1)
}

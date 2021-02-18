package p0100sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
}

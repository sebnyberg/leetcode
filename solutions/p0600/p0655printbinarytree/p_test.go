package p0655printbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return find(root, m, k)
}

func find(cur *TreeNode, m map[int]struct{}, k int) bool {
	if cur == nil {
		return false
	}
	if _, exists := m[k-cur.Val]; exists {
		return true
	}
	m[cur.Val] = struct{}{}
	return find(cur.Left, m, k) || find(cur.Right, m, k)
}

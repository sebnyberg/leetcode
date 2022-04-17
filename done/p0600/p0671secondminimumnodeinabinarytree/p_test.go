package p0671secondminimumnodeinabinarytree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	m := make(map[int]struct{})
	visit(root, m)
	l := make([]int, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	sort.Ints(l)
	if len(l) < 2 {
		return -1
	}
	return l[1]
}

func visit(cur *TreeNode, m map[int]struct{}) {
	if cur == nil {
		return
	}
	m[cur.Val] = struct{}{}
	visit(cur.Left, m)
	visit(cur.Right, m)
}

package p0606constructstringfrombinarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var res []byte
	dfs(root, &res)
	return string(res)
}

func dfs(cur *TreeNode, res *[]byte) {
	*res = append(*res, fmt.Sprint(cur.Val)...)
	if cur.Left != nil {
		*res = append(*res, '(')
		dfs(cur.Left, res)
		*res = append(*res, ')')
	}
	if cur.Right != nil {
		if cur.Left == nil {
			*res = append(*res, '(', ')')
		}
		*res = append(*res, '(')
		dfs(cur.Right, res)
		*res = append(*res, ')')
	}
}

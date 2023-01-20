package p3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := make(map[int]bool)
	for _, x := range to_delete {
		m[x] = true
	}
	var res []*TreeNode
	dfs(root, m, true, &res)
	return res
}

func dfs(curr *TreeNode, m map[int]bool, maybeRoot bool, res *[]*TreeNode) *TreeNode {
	if curr == nil {
		return nil
	}
	if m[curr.Val] {
		curr.Left = dfs(curr.Left, m, true, res)
		curr.Right = dfs(curr.Right, m, true, res)
		return nil
	}
	curr.Left = dfs(curr.Left, m, false, res)
	curr.Right = dfs(curr.Right, m, false, res)
	if maybeRoot {
		*res = append(*res, curr)
	}
	return curr
}

package p0102bstordertrav

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var t TreeTraverser
	t.traverse(root, 0)
	return t.res
}

type TreeTraverser struct {
	res [][]int
}

func (t *TreeTraverser) traverse(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(t.res) <= level {
		t.res = append(t.res, make([]int, 0))
	}
	t.res[level] = append(t.res[level], node.Val)
	t.traverse(node.Left, level+1)
	t.traverse(node.Right, level+1)
}

package p0107bstleveltrav

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var t TreeTraverser
	t.traverse(root, 1)
	for i, j := 0, len(t.res)-1; i < j; i, j = i+1, j-1 {
		t.res[i], t.res[j] = t.res[j], t.res[i]
	}
	return t.res
}

type TreeTraverser struct {
	res [][]int
}

func (t *TreeTraverser) traverse(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(t.res) < level {
		t.res = append(t.res, make([]int, 0))
	}
	t.res[level-1] = append(t.res[level-1], node.Val)
	t.traverse(node.Left, level+1)
	t.traverse(node.Right, level+1)
}

package p0103binzigzagtrav

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var t TreeTraverser
	t.traverse(root, 0)

	for i := 0; i < len(t.res); i++ {
		if i%2 == 1 {
			for j, k := 0, len(t.res[i])-1; j < k; j, k = j+1, k-1 {
				t.res[i][j], t.res[i][k] = t.res[i][k], t.res[i][j]
			}
		}
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
	if len(t.res) <= level {
		t.res = append(t.res, make([]int, 0))
	}
	t.res[level] = append(t.res[level], node.Val)
	t.traverse(node.Left, level+1)
	t.traverse(node.Right, level+1)
}

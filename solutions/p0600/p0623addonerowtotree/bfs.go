package p0623addonerowtotree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {

	dummy := &TreeNode{Left: root}
	curr := []*TreeNode{dummy}
	next := []*TreeNode{}

	for i := 2; i < depth; i++ {
		next = next[:0]
		for _, node := range curr {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		curr, next = next, curr
	}

	for _, node := range curr {
		node.Left = &TreeNode{Val: val, Left: node.Left}
		node.Right = &TreeNode{Val: val, Left: node.Right}
	}

	return dummy.Left
}

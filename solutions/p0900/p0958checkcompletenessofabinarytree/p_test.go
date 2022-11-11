package p0958checkcompletenessofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	// We can traverse the tree in level order using bfs and once a node is nil,
	// all following nodes must also be nil
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	var mustnil bool
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			if x == nil {
				mustnil = true
				continue
			}
			if mustnil {
				return false
			}
			next = append(next, x.Left, x.Right)
		}
		curr, next = next, curr
	}
	return true
}

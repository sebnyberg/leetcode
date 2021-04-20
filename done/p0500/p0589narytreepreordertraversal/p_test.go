package p0589narytreepreordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var out []int
	out = append(out, root.Val)
	for _, child := range root.Children {
		out = append(out, preorder(child)...)
	}
	return out
}

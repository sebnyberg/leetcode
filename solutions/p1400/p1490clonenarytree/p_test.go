package p1490clonenarytree

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	n := &Node{Val: root.Val}
	for _, ch := range root.Children {
		n.Children = append(n.Children, cloneTree(ch))
	}
	return n
}

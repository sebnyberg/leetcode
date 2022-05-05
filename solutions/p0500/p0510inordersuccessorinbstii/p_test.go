package p0510inordersuccessorinbstii

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func inorderSuccessor(node *Node) *Node {
	// There are only two cases:
	// Either the successor is the parent,
	// or it is the result of going right, [left, left, left, ...]
	res := node.Right
	if res == nil {
		p := node.Parent
		for p != nil {
			if p.Val > node.Val {
				return p
			}
			p = p.Parent
		}
		return nil
	}
	for res.Left != nil {
		res = res.Left
	}
	return res
}

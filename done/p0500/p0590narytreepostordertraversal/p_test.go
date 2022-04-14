package p0590narytreepostordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	collect(root, &res)
	return res
}

func collect(cur *Node, res *[]int) {
	if cur == nil {
		return
	}
	for _, child := range cur.Children {
		collect(child, res)
	}
	*res = append(*res, cur.Val)
}

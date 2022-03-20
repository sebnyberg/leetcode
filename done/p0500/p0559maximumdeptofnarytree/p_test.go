package p0559maximumdeptofnarytree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	curr := []*Node{root}
	next := []*Node{}
	level := 0
	for len(curr) > 0 {
		next = next[:0]
		for _, n := range curr {
			for _, ch := range n.Children {
				if ch != nil {
					next = append(next, ch)
				}
			}
		}
		curr, next = next, curr
		level++
	}
	return level
}

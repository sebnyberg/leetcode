package p0429narytreelevelordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	cur := []*Node{root}
	next := []*Node{}
	res := make([][]int, 0)
	var level int
	for len(cur) > 0 {
		next = next[:0]
		res = append(res, []int{})
		for _, n := range cur {
			res[level] = append(res[level], n.Val)
			next = append(next, n.Children...)
		}
		cur, next = next, cur
		level++
	}
	return res
}

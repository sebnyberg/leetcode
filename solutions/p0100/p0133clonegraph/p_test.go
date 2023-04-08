package p0133clonegraph

import "testing"

func TestCloneGraph(t *testing.T) {
	first := &Node{Val: 1}
	second := &Node{Val: 2}
	third := &Node{Val: 3}
	fourth := &Node{Val: 4}
	first.Neighbors = append(first.Neighbors, second, third)
	second.Neighbors = append(second.Neighbors, first, fourth)
	third.Neighbors = append(third.Neighbors, first, fourth)
	fourth.Neighbors = append(fourth.Neighbors, second, third)
	res := cloneGraph(first)
	_ = res

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodes := make(map[int]*Node)
	curr := []*Node{node}
	next := []*Node{}
	nodes[node.Val] = &Node{Val: node.Val}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, y := range x.Neighbors {
				if _, exists := nodes[y.Val]; !exists {
					nodes[y.Val] = &Node{Val: y.Val}
					next = append(next, y)
				}
				nodes[x.Val].Neighbors = append(nodes[x.Val].Neighbors, nodes[y.Val])
			}
		}
		curr, next = next, curr
	}

	return nodes[node.Val]
}

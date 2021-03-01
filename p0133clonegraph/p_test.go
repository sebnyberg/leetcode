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
	todos := []*Node{node}
	nodes[node.Val] = &Node{Val: node.Val}
	for len(todos) > 0 {
		var newTodo []*Node
		for _, todo := range todos {
			for _, neighbourNode := range todo.Neighbors {
				if _, exists := nodes[neighbourNode.Val]; !exists {
					nodes[neighbourNode.Val] = &Node{Val: neighbourNode.Val}
					newTodo = append(newTodo, neighbourNode)
				}
				nodes[todo.Val].Neighbors = append(nodes[todo.Val].Neighbors, nodes[neighbourNode.Val])
			}
		}

		todos = newTodo
	}

	return nodes[node.Val]
}

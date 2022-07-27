package leetcode

import (
	"strconv"
	"strings"
)

func (n *Node) String() string {
	// Leetcode spec
	//
	// For each visited node, add all children to the queue, then add null
	// At the end, strip trailing nulls (if any)
	//
	if n == nil {
		return "[null]"
	}
	s := func(x int) string {
		return strconv.Itoa(x)
	}
	nodes := []string{s(n.Val), "null"}

	queue := []*Node{n}
	for i := 0; i < len(queue); i++ {
		cur := queue[i]
		for _, child := range cur.Children {
			nodes = append(nodes, s(child.Val))
			queue = append(queue, child)
		}
		nodes = append(nodes, "null")
	}
	// Strip trailing nulls
	var i int
	for i = len(nodes); i > 0 && nodes[i-1] == "null"; i-- {
	}
	nodes = nodes[:i]
	return "[" + strings.Join(nodes, ",") + "]"
}

func ParseNaryTree(tree string) *Node {
	strs := strings.Split(tree[1:len(tree)-1], ",")

	// Create dummy node in first position
	nodes := []*Node{{Val: -1}}

	queue := []int{0}
	for i, j := 0, 0; i < len(strs) && j < len(queue); i, j = i+1, j+1 {
		cur := nodes[queue[j]]
		for i < len(strs) && strs[i] != "null" {
			// Add child to current node and to the queue
			v, _ := strconv.Atoi(strs[i])
			nodes = append(nodes, &Node{Val: v})
			idx := len(nodes) - 1
			cur.Children = append(cur.Children, nodes[idx])
			queue = append(queue, idx)
			i++
		}
	}
	if len(nodes[0].Children) == 0 {
		return nil
	}
	return nodes[0].Children[0]
}

type Node struct {
	Val      int
	Children []*Node
}

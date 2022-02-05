package p0428serializeanddeserializenarytree

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Codec(t *testing.T) {
	tree := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2},
			{
				Val: 3,
				Children: []*Node{
					{Val: 6},
					{
						Val: 7,
						Children: []*Node{
							{
								Val:      11,
								Children: []*Node{{Val: 14}},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val:      8,
						Children: []*Node{{Val: 12}},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val:      9,
						Children: []*Node{{Val: 13}},
					},
					{Val: 10},
				},
			},
		},
	}
	enc := Constructor()
	serialized := enc.serialize(tree)
	want := "[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]"
	require.Equal(t, want, serialized)
	res2 := enc.deserialize(serialized)
	_ = res2
	require.Equal(t, "[null]", enc.serialize(nil))
}

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	// Leetcode spec
	//
	// For each visited node, add all children to the queue, then add null
	// At the end, strip trailing nulls (if any)
	//
	if root == nil {
		return "[null]"
	}
	s := func(x int) string {
		return strconv.Itoa(x)
	}
	nodes := []string{s(root.Val), "null"}

	queue := []*Node{root}
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

func (this *Codec) deserialize(data string) *Node {
	strs := strings.Split(data[1:len(data)-1], ",")

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

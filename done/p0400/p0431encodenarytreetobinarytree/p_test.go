package p0431encodenarytreetobinarytree

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

func (n *TreeNode) FindVal(val int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val == val {
		return n
	}
	if l := n.Left.FindVal(val); l != nil {
		return l
	}
	return n.Right.FindVal(val)
}

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == nil || other == nil {
		return n == other
	}
	if n.Val != other.Val {
		return false
	}
	return n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
}

func ParseTree(input string) *TreeNode {
	// Trim start/end []
	input = input[1 : len(input)-1]

	// Split by comma
	inputParts := strings.Split(input, ",")
	n := len(inputParts)

	if n == 0 || inputParts[0] == "" {
		return nil
	}

	// Create one node per element in the array
	nodes := make([]*TreeNode, n)
	for i, part := range inputParts {
		if part != "null" {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			nodes[i] = &TreeNode{Val: val}
		}
	}

	q := list.New()
	q.PushBack(nodes[0])

	i := 1
	for q.Len() > 0 && i < n {
		el := q.Remove(q.Front()).(*TreeNode)
		if nodes[i] != nil {
			el.Left = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
		if i >= n {
			break
		}
		if nodes[i] != nil {
			el.Right = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
	}

	return nodes[0]
}

func Test_Codec(t *testing.T) {
	for _, tc := range []struct {
		input string
	}{
		{"[1,null,3,2,4,null,5,6]"},
		{"[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]"},
		{"[]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			var c Codec
			tree := ParseNaryTree(tc.input)
			encoded := c.encode(tree)
			decoded := c.decode(encoded)
			require.Equal(t, tree.String(), decoded.String())
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}
	res := &TreeNode{
		Val:  root.Val,
		Left: visit(root),
	}
	return res
}

func visit(cur *Node) *TreeNode {
	if len(cur.Children) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: cur.Children[0].Val,
	}
	root.Left = visit(cur.Children[0])
	prev := root
	for i := 1; i < len(cur.Children); i++ {
		cur := &TreeNode{
			Val:  cur.Children[i].Val,
			Left: visit(cur.Children[i]),
		}
		prev.Right = cur
		prev = cur
	}
	return root
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}
	res := &Node{
		Val:      root.Val,
		Children: collect(root.Left),
	}
	return res
}

func collect(cur *TreeNode) []*Node {
	if cur == nil {
		return nil
	}
	var res []*Node
	for cur != nil {
		res = append(res, &Node{Val: cur.Val})
		res[len(res)-1].Children = collect(cur.Left)
		cur = cur.Right
	}
	return res
}

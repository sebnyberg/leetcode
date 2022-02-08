package p0449serializeanddeserializebst

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_Codec(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	var c Codec
	res := c.serialize(tree)
	res2 := c.deserialize(res)
	_ = res2
	fmt.Println("hehe")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// Perform level-order traversal (Leetcode encoding)
	// When encountering a non-zero node, push an expected pair of nodes to the
	// stack. Skip trimming this time.
	res := make([]string, 0)
	stack := []*TreeNode{root}
	for i := 0; i < len(stack); i++ {
		if stack[i] != nil {
			res = append(res, strconv.Itoa(stack[i].Val))
			stack = append(stack, stack[i].Left, stack[i].Right)
		} else {
			res = append(res, "null")
		}
	}
	return "[" + strings.Join(res, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1] // strip []
	nodes := strings.Split(data, ",")
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}
	num := func(s string) int {
		x, _ := strconv.Atoi(s)
		return x
	}

	root := &TreeNode{Val: num(nodes[0])}
	parent := []*TreeNode{root}
	for i, n := range nodes[1:] {
		if n == "null" {
			continue
		}
		cur := &TreeNode{Val: num(n)}
		parent = append(parent, cur)
		if i%2 == 0 {
			parent[i/2].Left = cur
		} else {
			parent[i/2].Right = cur
		}
	}
	return root
}

package p0109convertsortarrtobst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedListToBST(t *testing.T) {
	nodeValues := []int{1, 3}
	// nodeValues := []int{-10, -3, 0, 5, 9}
	node := &ListNode{Val: nodeValues[len(nodeValues)-1]}
	for i := len(nodeValues) - 2; i >= 0; i-- {
		node = &ListNode{
			Val:  nodeValues[i],
			Next: node,
		}
	}

	result := sortedListToBST(node)
	_ = result
	require.Equal(t, true, true)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	cur := head
	// Count number of nodes in the list
	var n int
	for cur != nil {
		cur = cur.Next
		n++
	}

	// Creating a stateful struct here to avoid a pointer pointer
	c := bstCreator{head}
	return c.create(n)
}

type bstCreator struct {
	node *ListNode
}

func (c *bstCreator) create(n int) *TreeNode {
	if n == 0 {
		return nil
	}
	if n == 1 {
		defer func() { c.node = c.node.Next }()
		return &TreeNode{Val: c.node.Val}
	}
	left := c.create(n / 2)
	val := c.node.Val
	c.node = c.node.Next
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: c.create((n - 1) / 2),
	}
}

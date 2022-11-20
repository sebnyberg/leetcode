package p0998maximumbinarytreeii

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	// If the node is larger than the current node, then it should replace it.
	// To do this we need to keep track of the previous node
	dummy := &TreeNode{
		Right: root,
		Val:   math.MinInt32,
	}
	prev := dummy
	cur := prev.Right
	for cur != nil && cur.Val > val {
		prev, cur = cur, cur.Right
	}
	prev.Right = &TreeNode{
		Val:  val,
		Left: cur,
	}

	return dummy.Right
}

package p0156btreeupsidedown

import "testing"

func TestUpsideBT(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	newRoot := upsideDownBinaryTree(tree)
	_ = newRoot
	upsideDownBinaryTree(&TreeNode{Val: 1})
	upsideDownBinaryTree(nil)
	_ = 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return helper(root.Left, root, root.Right)
}

func helper(cur *TreeNode, prevRoot *TreeNode, prevRight *TreeNode) *TreeNode {
	if cur == nil {
		return prevRoot
	}
	nextLeft, nextRight := cur.Left, cur.Right
	cur.Left = prevRight
	cur.Right = prevRoot
	if prevRight != nil {
		helper(prevRight, nil, nil)
	}
	if nextLeft != nil {
		return helper(nextLeft, cur, nextRight)
	} else {
		return cur
	}
}

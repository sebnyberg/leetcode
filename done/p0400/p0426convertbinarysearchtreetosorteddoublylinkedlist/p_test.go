package p0426convertbinarysearchtreetosorteddoublylinkedlist

import "testing"

func Test_treeToDoublyList(t *testing.T) {
	root := &Node{
		Val: 4,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 3},
		},
		Right: &Node{Val: 5},
	}
	res := treeToDoublyList(root)
	_ = res
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := &Node{}
	pred := head
	visit(root, &pred)
	// The final node in the tree is currently stored in pred.
	// Pred.Right is missing, and the first nodes left pointer is going to head
	first := head.Right
	pred.Right = first
	first.Left = pred
	return first
}

func visit(cur *Node, pred **Node) {
	if cur == nil {
		return
	}
	visit(cur.Left, pred)
	right := cur.Right
	(*pred).Right, cur.Left, *pred = cur, *pred, cur
	visit(right, pred)
}

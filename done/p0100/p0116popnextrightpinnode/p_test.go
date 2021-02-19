package p0116popnextrightpinnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectPair(root.Left, root.Right)
	return root
}

func connectPair(left *Node, right *Node) {
	if left == nil {
		return
	}
	left.Next = right
	connectPair(left.Left, left.Right)
	connectPair(left.Right, right.Left)
	connectPair(right.Left, right.Right)
}

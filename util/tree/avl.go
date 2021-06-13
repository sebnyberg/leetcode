package tree

// AVLTree is a tree where the absolute difference of the heights of the left
// and right child-trees is <= 1.
type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Insert(x int) {
	if t.root == nil {
		t.root = &AVLNode{
			Val: x,
			h:   1,
		}
		return
	}
	t.root = t.root.Insert(x)
}

type AVLNode struct {
	Val   int
	Left  *AVLNode
	Right *AVLNode
	h     int
}

func (t *AVLNode) Height() int {
	if t == nil {
		return 0
	}
	return t.h
}

// Insert inserts the provided value and updates the height of this tree.
// Calling this function may cause a rebalance operation, in which case
// the returned value is the new root of the sub-tree. Ensure that the caller
// updates references accordingly.
func (n *AVLNode) Insert(x int) *AVLNode {
	if x < n.Val {
		if n.Left == nil {
			n.Left = &AVLNode{Val: x, h: 1}
		} else {
			n.Left = n.Left.Insert(x)
		}
	} else {
		if n.Right == nil {
			n.Right = &AVLNode{Val: x, h: 1}
		} else {
			n.Right = n.Right.Insert(x)
		}
	}
	n.updateHeight()

	balance := n.Right.Height() - n.Left.Height()

	if balance < -1 {
		// LL
		if n.Left.Left.Height() > n.Left.Right.Height() {
			return n.rotateRight()
		} else { // LR
			n.Left = n.Left.rotateLeft()
			return n.rotateRight()
		}
	} else if balance > 1 {
		// RR
		if n.Left.Left.Height() > n.Left.Right.Height() {
			return n.rotateLeft()
		} else { // RL
			n.Right = n.Right.rotateRight()
			return n.rotateLeft()
		}
	}

	return n
}

func (n *AVLNode) updateHeight() {
	if n != nil {
		n.h = 1 + max(n.Left.Height(), n.Right.Height())
	}
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.Right
	n.Right.Left, n.Right = n, n.Right.Left
	newRoot.updateHeight()
	n.updateHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.Left
	n.Left.Right, n.Left = n, n.Left.Right
	newRoot.updateHeight()
	n.updateHeight()
	return newRoot
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

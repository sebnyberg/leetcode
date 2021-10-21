package tree

// AVLTree is a tree where the absolute difference of the heights of the left
// and right child-trees is <= 1.
type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) insert(x int) {
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
	left  *AVLNode
	right *AVLNode
	h     int
}

func (t *AVLNode) height() int {
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
		if n.left == nil {
			n.left = &AVLNode{Val: x, h: 1}
		} else {
			n.left = n.left.Insert(x)
		}
	} else {
		if n.right == nil {
			n.right = &AVLNode{Val: x, h: 1}
		} else {
			n.right = n.right.Insert(x)
		}
	}
	n.updateHeight()

	balance := n.right.height() - n.left.height()

	if balance < -1 {
		// LL
		if n.left.left.height() > n.left.right.height() {
			return n.rotateRight()
		} else { // LR
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if balance > 1 {
		// RR
		if n.right.right.height() > n.right.left.height() {
			return n.rotateLeft()
		} else { // RL
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}

	return n
}

func (n *AVLNode) updateHeight() {
	if n != nil {
		n.h = 1 + max(n.left.height(), n.right.height())
	}
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right.left, n.right = n, n.right.left
	newRoot.updateHeight()
	n.updateHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left.right, n.left = n, n.left.right
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

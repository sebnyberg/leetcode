package tree

// type BinaryNode struct {
// 	left  *BinaryNode
// 	right *BinaryNode
// 	val   int
// }

// type BinaryTree struct {
// 	root *BinaryNode
// }

// func (t *BinaryTree) Insert(val int) {
// 	if t.root == nil {
// 		t.root = &BinaryNode{val: val}
// 	} else {
// 		t.root.Insert(val)
// 	}
// }

// func (n *BinaryNode) Insert(val int) {
// 	if n == nil {
// 		return
// 	}

// 	if val <= n.val {
// 		if n.left == nil {
// 			n.left = &BinaryNode{val: val}
// 		} else {
// 			n.left.Insert(val)
// 		}
// 		return
// 	}

// 	if n.right == nil {
// 		n.right = &BinaryNode{val: val}
// 	} else {
// 		n.right.Insert(val)
// 	}
// }

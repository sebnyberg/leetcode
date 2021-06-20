package tree

import "testing"

func TestAVLTree(t *testing.T) {
	var tree AVLTree
	tree.insert(40)
	tree.insert(20)
	tree.insert(10)
}

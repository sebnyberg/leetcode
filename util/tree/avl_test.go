package tree

import "testing"

func TestAVLTree(t *testing.T) {
	var tree AVLTree
	tree.Insert(40)
	tree.Insert(20)
	tree.Insert(10)
}

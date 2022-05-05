package p0173bstiterator

import "testing"

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
		Val:   2,
	}
	it := Constructor(root)
	for it.HasNext() {
		val := it.Next()
		_ = val
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur:   root,
		stack: make([]*TreeNode, 0),
	}
}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stack = append(this.stack, this.cur)
		this.cur = this.cur.Left
	}
	n := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1] // pop
	this.cur = n.Right
	return n.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.cur != nil
}

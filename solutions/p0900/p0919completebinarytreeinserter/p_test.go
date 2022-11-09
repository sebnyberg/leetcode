package p0919completebinarytreeinserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	i    int
	curr []*TreeNode
	next []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	var c CBTInserter
	c.root = root
	c.curr = []*TreeNode{root}
	c.next = []*TreeNode{}
	for {
		if c.i == len(c.curr) {
			c.curr, c.next = c.next, c.curr
			c.next = c.next[:0]
			c.i = 0
		}
		if c.curr[c.i].Left == nil {
			return c
		}
		if c.curr[c.i].Left != nil {
			c.next = append(c.next, c.curr[c.i].Left)
		}
		if c.curr[c.i].Right == nil {
			return c
		}
		if c.curr[c.i].Right != nil {
			c.next = append(c.next, c.curr[c.i].Right)
			c.i++
		}
	}
}

func (this *CBTInserter) Insert(val int) int {
	if this.i == len(this.curr) {
		this.curr, this.next = this.next, this.curr
		this.next = this.next[:0]
		this.i = 0
	}
	x := this.curr[this.i]
	m := &TreeNode{
		Val: val,
	}
	this.next = append(this.next, m)
	if x.Left == nil {
		x.Left = m
	} else {
		x.Right = m
		this.i++
	}
	return x.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

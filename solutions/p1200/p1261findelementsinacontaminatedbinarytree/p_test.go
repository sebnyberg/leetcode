package p1261findelementsinacontaminatedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	m map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	m := make(map[int]struct{})
	restore(m, root, 0)
	return FindElements{
		m: m,
	}
}

func restore(m map[int]struct{}, curr *TreeNode, val int) {
	if curr == nil {
		return
	}
	m[val] = struct{}{}
	curr.Val = val
	restore(m, curr.Left, val*2+1)
	restore(m, curr.Right, val*2+2)
}

func (this *FindElements) Find(target int) bool {
	_, exists := this.m[target]
	return exists
}

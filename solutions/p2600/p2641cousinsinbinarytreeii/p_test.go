package p2641cousinsinbinarytreeii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	root.Val = 0
	maybe := func(x *TreeNode) int {
		if x != nil {
			return x.Val
		}
		return 0
	}
	for len(curr) > 0 {
		next = next[:0]
		var sum int
		for _, x := range curr {
			sum += maybe(x.Left)
			sum += maybe(x.Right)
		}
		for _, x := range curr {
			pair := maybe(x.Left) + maybe(x.Right)
			if x.Left != nil {
				x.Left.Val = sum - pair
				next = append(next, x.Left)
			}
			if x.Right != nil {
				x.Right.Val = sum - pair
				next = append(next, x.Right)
			}
		}
		curr, next = next, curr
	}
	return root
}

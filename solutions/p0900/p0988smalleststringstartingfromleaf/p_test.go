package p0988smalleststringstartingfromleaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var res string
	var stack []byte
	visit(root, &stack, &res)
	return res
}

func visit(cur *TreeNode, stack *[]byte, res *string) {
	if cur == nil {
		return
	}
	*stack = append(*stack, byte(cur.Val+'a'))
	if cur.Left == nil && cur.Right == nil {
		s := rev(string(*stack))
		if *res == "" || s < *res {
			*res = s
		}
	} else {
		if cur.Left != nil {
			visit(cur.Left, stack, res)
		}
		if cur.Right != nil {
			visit(cur.Right, stack, res)
		}
	}
	*stack = (*stack)[:len(*stack)-1]
}

func rev(s string) string {
	res := []byte(s)
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

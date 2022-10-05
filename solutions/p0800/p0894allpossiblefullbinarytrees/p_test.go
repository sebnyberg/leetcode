package p0894allpossiblefullbinarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n&1 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	var res []*TreeNode
	for l := 1; ; l += 2 {
		r := n - l - 1
		if r < 1 {
			break
		}
		leftTrees := allPossibleFBT(l)
		rightTrees := allPossibleFBT(r)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				res = append(res, &TreeNode{
					Val:   0,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return res
}

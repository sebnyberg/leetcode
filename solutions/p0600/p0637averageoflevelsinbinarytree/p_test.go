package p0637averageoflevelsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	var res []float64
	for len(curr) > 0 {
		next = next[:0]
		var sum int
		for _, n := range curr {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(curr)))
		curr, next = next, curr
	}
	return res
}

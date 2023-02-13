package p1123lowestcommonancestorofdeepestleaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// The idea here is to perform BFS until there are no more nodes, collecting
	// paths from the root to the current nodes.
	//
	// Once done, we find the latest node in the path (from the root) which is
	// in common for all the deepest nodes.
	//
	curr := [][]*TreeNode{{root}}
	next := [][]*TreeNode{}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			n := len(x)
			if x[n-1].Left != nil {
				cpy := make([]*TreeNode, n)
				copy(cpy, x)
				cpy = append(cpy, x[n-1].Left)
				next = append(next, cpy)
			}
			if x[n-1].Right != nil {
				cpy := make([]*TreeNode, n)
				copy(cpy, x)
				cpy = append(cpy, x[n-1].Right)
				next = append(next, cpy)
			}
		}
		curr, next = next, curr
	}
	last := next
	var res *TreeNode
	for i := range last[0] {
		ok := true
		for j := range last {
			if last[j][i] != last[0][i] {
				ok = false
				break
			}
		}
		if ok {
			res = last[0][i]
		} else {
			break
		}
	}
	return res
}

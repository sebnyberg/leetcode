package p2415reverseoddlevelsofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	curr := []**TreeNode{&root}
	next := []**TreeNode{}
	for level := 1; len(curr) > 0; level++ {
		next = next[:0]
		// Collect nodes on level
		for _, node := range curr {
			x := *node
			if x.Left != nil {
				next = append(next, &x.Left)
			}
			if x.Right != nil {
				next = append(next, &x.Right)
			}
		}
		// On odd levels, reverse order
		if level&1 == 1 {
			for l, r := 0, len(next)-1; l < r; l, r = l+1, r-1 {
				*next[l], *next[r] = *next[r], *next[l]
				(*next[l]).Left, (*next[r]).Left = (*next[r]).Left, (*next[l]).Left
				(*next[l]).Right, (*next[r]).Right = (*next[r]).Right, (*next[l]).Right
			}
		}
		curr, next = next, curr
	}
	return root
}

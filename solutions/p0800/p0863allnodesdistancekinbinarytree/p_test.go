package p0863allnodesdistancekinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// Perform BFS to collect parent nodes of each node
	parent := make(map[*TreeNode]*TreeNode)
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			if x.Left != nil {
				parent[x.Left] = x
				next = append(next, x.Left)
			}
			if x.Right != nil {
				parent[x.Right] = x
				next = append(next, x.Right)
			}
		}
		curr, next = next, curr
	}
	curr = curr[:0]
	curr = append(curr, target)
	seen := make(map[*TreeNode]struct{})
	ok := func(x *TreeNode) bool {
		if x == nil {
			return false
		}
		if _, exists := seen[x]; exists {
			return false
		}
		return true
	}
	seen[target] = struct{}{}
	for steps := 0; steps < k; steps++ {
		next = next[:0]
		for _, x := range curr {
			if ok(x.Left) {
				seen[x.Left] = struct{}{}
				next = append(next, x.Left)
			}
			if ok(x.Right) {
				seen[x.Right] = struct{}{}
				next = append(next, x.Right)
			}
			if par := parent[x]; ok(par) {
				seen[par] = struct{}{}
				next = append(next, par)
			}
		}
		curr, next = next, curr
	}
	res := make([]int, len(curr))
	for i := range curr {
		res[i] = curr[i].Val
	}
	return res
}

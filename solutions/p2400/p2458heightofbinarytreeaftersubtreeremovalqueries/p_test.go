package p2458heightofbinarytreeaftersubtreeremovalqueries

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	// We need to:
	//
	// 1. Map each node number to its node
	// 2. Calculate the sum of each tree rooted at each node
	// 3. Map each node to its parent

	var nodes []*TreeNode
	var parent []*TreeNode
	var sizes []int
	ensureLen := func(i int) {
		for len(nodes) <= i {
			nodes = append(nodes, nil)
			parent = append(parent, nil)
			sizes = append(sizes, 0)
		}
	}
	var visit func(above, node *TreeNode) int
	visit = func(above, node *TreeNode) int {
		if node == nil {
			return 0
		}
		ensureLen(node.Val)
		nodes[node.Val] = node
		parent[node.Val] = above
		sizes[node.Val] = 1 + max(visit(node, node.Left), visit(node, node.Right))
		return sizes[node.Val]
	}
	visit(nil, root)
	res := make([]int, len(queries))
	for i, q := range queries {
		cur := nodes[q]
		p := parent[q]
		var size int
		for p != nil {
			if p.Left == cur {
				if p.Right != nil {
					size = 1 + max(size, sizes[p.Right.Val])
				} else {
					size++
				}
			} else {
				if p.Left != nil {
					size = 1 + max(size, sizes[p.Left.Val])
				} else {
					size++
				}
			}
			cur, p = p, parent[p.Val]
		}
		res[i] = size - 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

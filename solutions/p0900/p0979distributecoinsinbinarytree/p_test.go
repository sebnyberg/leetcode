package p0979distributecoinsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	// This problem was much trickier than the hard one that follows..
	//
	// The goal is to find an invariant where we can consider only the root and
	// its subtrees, then continue into each subtree.
	//
	// So how many times is the root involved in a move?
	//
	// Well, we know that a subtree that has too many coins will need to move
	// coins out of that tree. Each of those coins will move AT LEAST to this
	// node. So the afferent (incoming) contribution of coins to this node is
	// guaranteed to be equal to the overflow of coins in that subtree.
	//
	// Then, let's consider any subtree that has too few coins. Each missing
	// coin must move through the root of the tree to reach that subtree, and so
	// the efferent (outgoing) contribution of the root to a tree with missing
	// coins is equal to the lack of coins in that subtree.
	//
	// This gives us the solution. The change in moves contributed to the total
	// by the root of a tree is equal to the difference in coins vs nodes for
	// each subtree.
	var res int
	visit(root, &res)
	return res
}

func visit(cur *TreeNode, res *int) (nnodes, ncoins int) {
	if cur == nil {
		return 0, 0
	}
	leftNodes, leftCoins := visit(cur.Left, res)
	rightNodes, rightCoins := visit(cur.Right, res)
	leftCost := abs(leftNodes - leftCoins)
	rightCost := abs(rightNodes - rightCoins)
	*res += leftCost + rightCost
	return leftNodes + rightNodes + 1, leftCoins + rightCoins + cur.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

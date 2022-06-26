package p2322maximumscoreafterremovalsonatree

import (
	"fmt"
	"leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumScore(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		edges [][]int
		want  int
	}{
		{[]int{5, 5, 2, 4, 4, 2}, leetcode.ParseMatrix("[[0,1],[1,2],[5,2],[4,3],[1,3]]"), 0},
		{[]int{1, 5, 5, 4, 11}, leetcode.ParseMatrix("[[0,1],[1,2],[1,3],[3,4]]"), 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumScore(tc.nums, tc.edges))
		})
	}
}

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)

	// Create bi-directional adjacency list
	bidiAdj := make([][]int, n)
	for _, e := range edges {
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		bidiAdj[e[0]] = append(bidiAdj[e[0]], e[1])
		bidiAdj[e[1]] = append(bidiAdj[e[1]], e[0])
	}

	// Convert into regular adjacency list
	curr := []int{0}
	next := []int{}
	seen := make([]bool, n)
	adj := make([][]int, n)
	seen[0] = true
	for len(curr) > 0 {
		next = next[:0]
		for _, i := range curr {
			for _, j := range bidiAdj[i] {
				if seen[j] {
					continue
				}
				seen[j] = true
				next = append(next, j)
				adj[i] = append(adj[i], j)
			}
		}
		curr, next = next, curr
	}

	// Capture the XOR value of each subtree rooted in a node
	xors := make([]int, n)
	copy(xors, nums)
	var calcXOR func(i int) int
	calcXOR = func(i int) int {
		for _, nei := range adj[i] {
			xors[i] ^= calcXOR(nei)
		}
		return xors[i]
	}
	calcXOR(0)

	// Consider each node in the graph, let's call it the "first node"
	//
	// Visit each child of the first node
	//
	// If we cut the subtree above the first node, and above the child of the
	// first node, then we get the following:
	//
	// XOR of the first node without the child subtree is xor[first] ^ xor[child]
	// XOR of the child is xor[child]
	// XOR of the rest of the tree is xor[0] ^ xor[first]
	//
	// Once all children have been visited, we may visit all other nodes in the
	// graph which have not been marked as a parent, or as a child of the first
	// node. Removal of the subtrees rooted in those nodes is easy to calculate:
	//
	// XOR of the first node is xor[first]
	// XOR of the second node is xor[second]
	// XOR of the rest of the tree is xor[0] ^ xor[first] ^ xor[second]
	//

	// isParent[i] is true if the node is a isParent of the current node
	isParent := make([]bool, n)
	isParent[0] = true
	// isChild[i] is true if the node is a isChild of the current node
	// isChild is reset once traversal of all children for a node has finished
	isChild := make([]bool, n)

	res := math.MaxInt64

	// Visit each subchild, marking it as a child and comparing it with its
	// subtree parent.
	var visitSubChild func(subRoot, child int)
	visitSubChild = func(subRoot, child int) {
		subRootXOR := xors[subRoot] ^ xors[child]
		childXOR := xors[child]
		treeXOR := xors[0] ^ xors[subRoot]
		maxXOR := max(subRootXOR, max(childXOR, treeXOR))
		minXOR := min(subRootXOR, min(childXOR, treeXOR))
		res = min(res, maxXOR-minXOR)
		isChild[child] = true
		for _, nei := range adj[child] {
			visitSubChild(subRoot, nei)
		}
	}

	var visitSubRoot func(subRoot int)
	visitSubRoot = func(subRoot int) {
		// Collect children for this node and evaluate cutting the two subtrees
		// rooted in subRoot and the child.
		for i := range isChild {
			isChild[i] = false
		}
		for _, child := range adj[subRoot] {
			visitSubChild(subRoot, child)
		}
		isParent[subRoot] = true

		// XOR with every node that is not a child or parent of the current one
		for other := 1; other < len(xors); other++ {
			if isParent[other] || isChild[other] {
				continue
			}
			a := xors[other]
			b := xors[subRoot]
			c := xors[0] ^ a ^ b
			res = min(res, max(a, max(b, c))-min(a, min(b, c)))
		}

		// Consider each child as the root of a subtree
		for _, child := range adj[subRoot] {
			visitSubRoot(child)
		}
		isParent[subRoot] = false
	}

	for _, subRoot := range adj[0] {
		visitSubRoot(subRoot)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2925maximumscoreafterapplyingoperationsonatree

import (
	"math"
)

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	// We can revert the problem definition and say: we want to pick one node
	// per path from root to leaf and remove that from the final sum. How do we
	// minimize the total sum of removed nodes?
	//
	// For any non-zero node, we can choose either to remove that node and be
	// done with the subtree, or remove one node from both the left and right
	// trees.

	n := len(values)
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	seen := make([]bool, n)

	var sum int
	for i := range values {
		sum += values[i]
	}

	res := minRemoval(seen, adj, values, 0)
	return int64(sum - res)
}

func minRemoval(seen []bool, adj [][]int, values []int, i int) int {
	seen[i] = true

	res := math.MaxInt64
	if values[i] != 0 {
		res = values[i]
	}
	var childSum int
	var adjCount int
	for _, x := range adj[i] {
		if seen[x] {
			continue
		}
		adjCount++
		a := minRemoval(seen, adj, values, x)
		if a == math.MaxInt64 {
			// Cannot skip current node - one of the subtrees cannot have a
			// removal within each path
			childSum = math.MaxInt64
			break
		}
		childSum += a
	}
	if adjCount == 0 {
		childSum = math.MaxInt64
	}

	return min(res, childSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0834sumofdistancesintree

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{6, leetcode.ParseMatrix("[[0,1],[0,2],[2,3],[2,4],[2,5]]"), []int{8, 12, 6, 10, 10, 10}},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{1, 0}}, []int{1, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, sumOfDistancesInTree(tc.n, tc.edges))
		})
	}
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// Any node can be picked as a root node of the tree.
	// We can use this to calculate the number of nodes and sum of distances for a
	// sub-tree.
	type result struct {
		count, sum int
	}
	subRes := make([]result, n)
	var nilResult result
	var dfs func(i int) result
	dfs = func(i int) result {
		if subRes[i] != nilResult {
			return result{}
		}
		subRes[i].count = 1
		for _, nei := range adj[i] {
			r := dfs(nei)
			subRes[i].count += r.count
			subRes[i].sum += r.count + r.sum
		}
		return subRes[i]
	}
	dfs(0)

	// Now we can traverse a second time to calculate the exact sum of distances
	// for each node. If we dfs just like before, the order of visits will be
	// correct.
	seen := make([]bool, n)
	var dfs2 func(i int, parentRes result)
	dfs2 = func(i int, parentRes result) {
		// The result of the current node is combined with the parent result to form
		// the actual result.
		subRes[i].count += parentRes.count
		subRes[i].sum += parentRes.sum

		// The contribution of this node to its children will be the difference
		// between the child's result and this node's result.
		for _, nei := range adj[i] {
			if seen[nei] {
				continue
			}
			seen[nei] = true
			sumContrib := subRes[nei].sum + subRes[nei].count
			parentCount := subRes[i].count - subRes[nei].count
			parentSum := subRes[i].sum - sumContrib + parentCount
			parentRes := result{
				count: parentCount,
				sum:   parentSum,
			}
			dfs2(nei, parentRes)
		}
	}
	seen[0] = true
	dfs2(0, result{})
	res := make([]int, n)
	for i := range res {
		res[i] = subRes[i].sum
	}

	return res
}

package p2603collectcoinsinatree

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_collectTheCoins(t *testing.T) {
	for i, tc := range []struct {
		coins []int
		edges [][]int
		want  int
	}{
		{
			[]int{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1},
			leetcode.ParseMatrix("[[0,1],[1,2],[2,3],[2,4],[2,5],[2,6],[4,7],[6,8],[5,9],[4,10],[6,11]]"),
			0,
		},
		{
			[]int{0, 1, 0},
			leetcode.ParseMatrix("[[0,1],[0,2]]"),
			0,
		},
		{
			[]int{0, 0, 0},
			leetcode.ParseMatrix("[[0,1],[0,2]]"),
			0,
		},
		{
			[]int{0, 0, 0, 1, 1, 0, 0, 1},
			leetcode.ParseMatrix("[[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]"),
			2,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, collectTheCoins(tc.coins, tc.edges))
		})
	}
}

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)

	// Use adj map for easier path trimming
	adj := make([]map[int]bool, n)
	for i := range adj {
		adj[i] = make(map[int]bool)
	}
	indeg := make([]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a][b] = true
		adj[b][a] = true
		indeg[a]++
		indeg[b]++
	}
	leaves := []int{}
	for j, deg := range indeg {
		if deg == 1 {
			leaves = append(leaves, j)
		}
	}

	// Remove any leaves which do not have coins
	next := []int{}
	nextLeaves := []int{}
	for len(leaves) > 0 {
		next = next[:0]
		for _, x := range leaves {
			if coins[x] == 1 {
				nextLeaves = append(nextLeaves, x)
				continue
			}
			// This node must be removed
			for y := range adj[x] {
				delete(adj[y], x)
				indeg[y]--
				if indeg[y] == 1 {
					next = append(next, y)
				}
			}
		}
		leaves, next = next, leaves
	}

	// Perform two "free" rounds of trimming leaves from the tree
	leaves = nextLeaves
	for k := 0; k < 2; k++ {
		next = next[:0]
		for _, x := range leaves {
			for y := range adj[x] {
				delete(adj[y], x)
				indeg[y]--
				if indeg[y] == 1 {
					next = append(next, y)
				}
			}
		}
		leaves, next = next, leaves
	}

	// If there is <= 1 node that needs to be visited, simply start in that
	// position.
	if len(leaves) <= 1 {
		return 0
	}

	// Now, we have a graph where every leaf node must be visited once.
	// It does not matter where we start - the result will always be the same.
	// Perform DFS to calculate the path length
	seen := make([]bool, n)
	seen[leaves[0]] = true
	res := dfs(seen, adj, leaves[0])

	return res
}

func dfs(seen []bool, adj []map[int]bool, x int) int {
	var res int
	for y := range adj[x] {
		if seen[y] {
			continue
		}
		seen[y] = true
		res += 2 + dfs(seen, adj, y)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2646minimizethetotalpriceofthetrips

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumTotalPrice(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
		want  int
	}{
		{
			4,
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3]]"),
			[]int{2, 2, 10, 6},
			leetcode.ParseMatrix("[[0,3],[2,1],[2,3]]"),
			23,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTotalPrice(tc.n, tc.edges, tc.price, tc.trips))
		})
	}
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	adj := make([][]int, n)
	for _, x := range edges {
		a := x[0]
		b := x[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Because this is a tree, there are never multiple possible paths for a
	// given trip. In other words, we may first take all trips, counting how
	// often we end up at each node, then we find a way to select which nodes
	// would minimize the total price sum for all trips.
	seen := make([]bool, n)
	visits := make([]int, n)
	for _, t := range trips {
		from := t[0]
		to := t[1]
		visits[from]++
		seen[from] = true
		travel(seen, adj, visits, from, to)
		seen[from] = false
	}

	// Find a leaf
	leaf := -1
	for i := range adj {
		if len(adj[i]) == 1 {
			leaf = i
			break
		}
	}
	if leaf == -1 {
		leaf = 0
	}

	// Now use memoized dfs to try all possible combinations of zeroing
	mem := make(map[key]int)
	seen[leaf] = true
	res := dfs(seen, mem, adj, visits, price, leaf, true)

	return res
}

type key struct {
	i      int
	zeroed bool
}

// dfs finds the lowest price sum for the unvisited tree given that the current
// node is zeroed or not
func dfs(seen []bool, mem map[key]int, adj [][]int, visits, price []int, i int, canHalf bool) int {
	k := key{i, canHalf}
	if v, exists := mem[k]; exists {
		return v
	}

	// We may always choose to not half the price at this node
	res := visits[i] * price[i]
	for _, nei := range adj[i] {
		if seen[nei] {
			continue

		}
		seen[nei] = true
		res += dfs(seen, mem, adj, visits, price, nei, true)
		seen[nei] = false
	}

	// Or, if we can half this node, then not half the neighbours
	if canHalf {
		halfRes := (visits[i] * price[i]) / 2
		for _, nei := range adj[i] {
			if seen[nei] {
				continue
			}
			seen[nei] = true
			halfRes += dfs(seen, mem, adj, visits, price, nei, false)
			seen[nei] = false
		}
		if halfRes < res {
			res = halfRes
		}
	}

	mem[k] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// travel travels to the target storing visits along the way in visits.
func travel(seen []bool, adj [][]int, visits []int, i, target int) bool {
	if i == target {
		return true
	}
	for _, nei := range adj[i] {
		if seen[nei] {
			continue
		}
		seen[nei] = true
		visits[nei]++
		ok := travel(seen, adj, visits, nei, target)
		seen[nei] = false
		if ok {
			return true
		}
		visits[nei]--
	}
	return false
}

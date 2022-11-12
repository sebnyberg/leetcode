package p2467mostprofitablepathinatree

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_mostProfitablePath(t *testing.T) {
	for i, tc := range []struct {
		edges  [][]int
		bob    int
		amount []int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3],[3,4]]"),
			3,
			[]int{-2, 4, 2, -4, 6},
			6,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, mostProfitablePath(tc.edges, tc.bob, tc.amount))
		})
	}
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	seen := make([]bool, n)
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	seen[0] = true
	curr := []int{0}
	next := []int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			var j int
			for _, nei := range adj[x] {
				if seen[nei] {
					continue
				}
				seen[nei] = true
				next = append(next, nei)
				adj[x][j] = nei
				j++
			}
			adj[x] = adj[x][:j]
		}
		curr, next = next, curr
	}
	bobPath := []int{0}
	findBob(adj, &bobPath, 0, bob)
	for i := len(bobPath) - 1; i > len(bobPath)/2; i-- {
		amount[bobPath[i]] = 0
	}
	mid := len(bobPath) / 2
	if len(bobPath)&1 == 0 {
		amount[bobPath[mid]] = 0
	} else {
		amount[bobPath[mid]] /= 2
	}

	// Find optimal path for alice
	res := findOptimalPath(adj, amount, 0)
	return res
}

func findOptimalPath(adj [][]int, amount []int, curr int) int {
	res := amount[curr]
	subCost := math.MinInt64
	for _, nei := range adj[curr] {
		subCost = max(subCost, findOptimalPath(adj, amount, nei))
	}
	if subCost != math.MinInt64 {
		res += subCost
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findBob(adj [][]int, path *[]int, curr, bob int) bool {
	if curr == bob {
		return true
	}
	for _, child := range adj[curr] {
		*path = append(*path, child)
		if findBob(adj, path, child, bob) {
			return true
		}
		*path = (*path)[:len(*path)-1]
	}
	return false
}

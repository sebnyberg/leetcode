package p2508addedgestomakedegreesofallnodeseven

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_isPossible(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		want  bool
	}{
		{
			4,
			leetcode.ParseMatrix("[[1,2],[1,3],[1,4]]"),
			false,
		},
		{
			11,
			leetcode.ParseMatrix("[[5,9],[8,1],[2,3],[7,10],[3,6],[6,7],[7,8],[5,1],[5,7],[10,11],[3,7],[6,11],[8,11],[3,4],[8,9],[9,1],[2,10],[9,11],[5,11],[2,5],[8,10],[2,7],[4,1],[3,10],[6,1],[4,9],[4,6],[4,5],[2,4],[2,11],[5,8],[6,9],[4,10],[3,11],[4,7],[3,5],[7,1],[2,9],[6,10],[10,1],[5,6],[3,9],[2,6],[7,9],[4,11],[4,8],[6,8],[3,8],[9,10],[5,10],[2,8],[7,11]]"),
			false,
		},
		{
			4,
			leetcode.ParseMatrix("[[4,1],[3,2],[2,4],[1,3]]"),
			false,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isPossible(tc.n, tc.edges))
		})
	}
}

func isPossible(n int, edges [][]int) bool {
	// Just count the number of uneven indegree nodes in the graph....
	indeg := make([]int, n)
	adj := make(map[int]map[int]bool)
	must := func(i int) {
		if _, exists := adj[i]; !exists {
			adj[i] = make(map[int]bool)
		}
	}
	for _, e := range edges {
		a, b := e[0]-1, e[1]-1
		indeg[a]++
		indeg[b]++
		must(a)
		must(b)
		adj[a][b] = true
		adj[b][a] = true
	}
	var odds []int
	var nodd int
	for i, d := range indeg {
		if d&1 == 1 {
			nodd++
			odds = append(odds, i)
		}
	}
	// Obviously, we need to connect odds.
	// If there's an odd number of odds, there is no solution.
	if nodd&1 == 1 || nodd > 4 {
		return false
	}
	if nodd == 0 {
		return true
	}

	// If nodd == 4, then we must connect odds to odds only. That means that
	// there must exist a valid pairing
	if nodd == 4 {
		return hasTwoPair(odds, adj)
	}

	// If nodd == 2, then we can either directly connect the two nodes, or find
	// an even node which does not connect to either of the odd nodes
	a := odds[0]
	b := odds[1]
	if !adj[a][b] {
		return true
	}
	// Find a node which isn't a or b and doesn't connect to either of them
	for node := range adj {
		if node == a || node == b || adj[node][a] || adj[node][b] {
			continue
		}
		return true
	}
	return false
}
func hasTwoPair(odds []int, adj map[int]map[int]bool) bool {
	var seen [4]bool
	for i := range odds {
		if seen[i] {
			continue
		}
		seen[i] = true
		var match bool
		for j := range odds {
			if i == j || seen[j] || adj[odds[i]][odds[j]] {
				continue
			}
			match = true
			seen[j] = true
			break
		}
		if !match {
			return false
		}
	}
	return true
}

func findSolution(odds []int, adj [][]int, i, bm int) bool {
	if i == len(odds) {
		return true
	}
	if bm&(1<<i) > 0 {
		return findSolution(odds, adj, i+1, bm)
	}
	bm |= (1 << i)
	for j := 0; j < len(odds); j++ {
		if bm&(1<<j) > 0 {
			continue
		}
		for _, x := range adj[odds[i]] {
			if x == odds[j] {
				goto skip
			}
		}
		// Try this pairing
		if findSolution(odds, adj, i+1, bm|(1<<j)) {
			return true
		}
	skip:
	}
	return false
}

package p2581countnumberofpossiblerootnodes

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_rootCount(t *testing.T) {
	for i, tc := range []struct {
		edges   [][]int
		guesses [][]int
		k       int
		want    int
	}{
		{
			leetcode.ParseMatrix("[[0,1],[1,2],[2,3],[3,4]]"),
			leetcode.ParseMatrix("[[1,0],[3,4],[2,1],[3,2]]"),
			1,
			5,
		},
		{
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3],[4,2]]"),
			leetcode.ParseMatrix("[[1,3],[0,1],[1,0],[2,4]]"),
			3,
			3,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, rootCount(tc.edges, tc.guesses, tc.k))
		})
	}
}

func rootCount(edges [][]int, guesses [][]int, k int) int {
	// For any given rooted tree, it is possible to calculate how many guesses
	// were right in O(n).
	//
	// Whenever wandering across an edge from guesses, the total number of valid
	// guesses change according to the direction of the edge.
	n := len(edges) + 1
	adj := make([][]int, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	deltas := make(map[[2]int]int)
	for _, g := range guesses {
		deltas[[2]int{g[0], g[1]}]++
	}
	seen := make([]bool, n)
	curr := []int{0}
	next := []int{}
	var nvalid int
	seen[0] = true
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, nei := range adj[x] {
				if seen[nei] {
					continue
				}
				seen[nei] = true
				if deltas[[2]int{x, nei}] > 0 {
					nvalid++
				}
				next = append(next, nei)
			}
		}
		curr, next = next, curr
	}

	var res int
	if nvalid >= k {
		res++
	}
	for i := range seen {
		seen[i] = false
	}
	for _, g := range guesses {
		deltas[[2]int{g[1], g[0]}]--
	}
	seen[0] = true
	curr2 := [][2]int{{0, nvalid}}
	next2 := [][2]int{}
	for len(curr2) > 0 {
		next2 = next2[:0]
		for _, x := range curr2 {
			u := x[0]
			nvalid := x[1]
			for _, nei := range adj[u] {
				if seen[nei] {
					continue
				}
				seen[nei] = true
				y := nvalid + deltas[[2]int{nei, u}]
				if y >= k {
					res++
				}
				next2 = append(next2, [2]int{nei, y})
			}
		}
		curr2, next2 = next2, curr2
	}
	return res
}

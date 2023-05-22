package p1377frogpositionaftertseconds

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_frogPosition(t *testing.T) {
	for i, tc := range []struct {
		n      int
		edges  [][]int
		t      int
		target int
		want   float64
	}{
		{
			7,
			leetcode.ParseMatrix("[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]"),
			2,
			4,
			1.0 / 6,
		},
		{
			7,
			leetcode.ParseMatrix("[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]"),
			1,
			7,
			1.0 / 3,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.InEpsilon(t, tc.want, frogPosition(tc.n, tc.edges, tc.t, tc.target), 0.005)
		})
	}
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// This might seem like a BFS problem but there is a need to keep track of
	// the prior path taken by the frog, otherwise we would not be able to
	// handle cycles in the graph.
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0] - 1
		b := e[1] - 1
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	target--
	var seen [101]bool
	seen[0] = true
	res := dfs(adj, &seen, 0, n, t, target)
	return res
}

func dfs(adj [][]int, seen *[101]bool, i, n, t, target int) float64 {
	if t == 0 {
		if i == target {
			return 1
		}
		return 0
	}
	var res float64
	var choices int
	for _, b := range adj[i] {
		if seen[b] {
			continue
		}
		seen[b] = true
		choices++
		res += dfs(adj, seen, b, n, t-1, target)
		seen[b] = false
	}
	if choices == 0 {
		return 0
	}
	return res / float64(choices)
}

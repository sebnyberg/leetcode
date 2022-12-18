package p2497maximumstarsumofagraph

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxStarSum(t *testing.T) {
	for i, tc := range []struct {
		vals  []int
		edges [][]int
		k     int
		want  int
	}{
		{
			[]int{1, 2, 3, 4, 10, -10, -20},
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3],[3,4],[3,5],[3,6]]"),
			2,
			16,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxStarSum(tc.vals, tc.edges, tc.k))
		})
	}
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	res := vals[0]
	for i := range vals {
		var nei []int
		for _, x := range adj[i] {
			nei = append(nei, vals[x])
		}
		sort.Slice(nei, func(i, j int) bool {
			return nei[i] > nei[j]
		})
		v := vals[i]
		for j := 0; j < min(len(nei), k); j++ {
			if nei[j] <= 0 {
				break
			}
			v += nei[j]
		}
		res = max(res, v)
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

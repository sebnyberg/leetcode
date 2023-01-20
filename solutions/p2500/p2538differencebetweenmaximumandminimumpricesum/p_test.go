package p2538differencebetweenmaximumandminimumpricesum

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxOutput(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		price []int
		want  int64
	}{
		{
			8,
			leetcode.ParseMatrix("[[1,7],[2,3],[4,0],[5,7],[6,3],[3,0],[0,7]]"),
			[]int{4, 5, 6, 2, 2, 7, 7, 8},
			21,
		},
		{
			6,
			leetcode.ParseMatrix("[[0,1],[1,2],[1,3],[3,4],[3,5]]"),
			[]int{9, 8, 7, 6, 10, 5},
			24,
		},
		{
			3,
			leetcode.ParseMatrix("[[0,1],[1,2]]"),
			[]int{1, 1, 1},
			2,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxOutput(tc.n, tc.edges, tc.price))
		})
	}
}

type path struct {
	origin int
	sum    int
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	// The minimum path is always just the node itself.
	// The maximum path must include two leaf nodes, or else it could be
	// extended to get a better result.
	// The goal is therefore to pair two "leaf paths" from a node such that
	// their origin is different, and the sum is maximized.
	// Note that you have to pass both paths with and without the leaf node,
	// otherwise you may miss an optimum.
	if n == 1 {
		return 0
	}
	if n == 2 {
		a := max(price[0], price[1])
		return int64(a)
	}
	adj := make([][]int, n)
	indeg := make([]int, n)
	for i := range edges {
		a := edges[i][0]
		b := edges[i][1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		indeg[a]++
		indeg[b]++
	}
	withLeaf := make([][]path, n)
	withoutLeaf := make([][]path, n)
	curr := []int{}
	next := []int{}
	isleaf := make([]bool, n)
	for i, d := range indeg {
		if d == 1 {
			isleaf[i] = true
		}
	}
	for i := range indeg {
		if isleaf[i] {
			indeg[i] = math.MaxInt32
			for _, nei := range adj[i] {
				withLeaf[nei] = append(withLeaf[nei], path{i, price[i]})
				withoutLeaf[nei] = append(withoutLeaf[nei], path{i, 0})
				indeg[nei]--
			}
		}
	}
	for i, d := range indeg {
		if d <= 1 {
			curr = append(curr, i)
		}
	}
	var res int64
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			indeg[x] = math.MaxInt32
			with := withLeaf[x]
			without := withoutLeaf[x]
			sort.Slice(with, func(i, j int) bool {
				return with[i].sum > with[j].sum
			})
			sort.Slice(without, func(i, j int) bool {
				return without[i].sum > without[j].sum
			})
			var best int
			if with[0].origin == without[0].origin {
				if len(with) > 1 {
					best = max(
						with[0].sum+without[1].sum,
						with[1].sum+without[0].sum,
					) + price[x]
				} else {
					best = price[x] + without[0].sum
				}
			} else {
				best = with[0].sum + without[0].sum + price[x]
			}
			bestWith := path{x, with[0].sum + price[x]}
			bestWithout := path{x, without[0].sum + price[x]}
			if int64(best) > res {
				res = int64(best)
			}
			for _, nei := range adj[x] {
				if indeg[nei] == math.MaxInt32 {
					continue
				}
				withLeaf[nei] = append(withLeaf[nei], bestWith)
				withoutLeaf[nei] = append(withoutLeaf[nei], bestWithout)
				indeg[nei]--
				if indeg[nei] == 1 {
					next = append(next, nei)
				}
			}
		}
		curr, next = next, curr
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

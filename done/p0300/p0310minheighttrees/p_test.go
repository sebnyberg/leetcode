package p0310minheighttrees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinHeightTrees(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{4, [][]int{{1, 0}, {1, 2}, {1, 3}}, []int{1}},
		{6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, []int{3, 4}},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{0, 1}}, []int{0, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findMinHeightTrees(tc.n, tc.edges))
		})
	}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	adj := make([][]int, n)
	nedges := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		nedges[a]++
		nedges[b]++
	}

	next := make([]int, 0, n)
	removed := make([]bool, n)
	for n > 2 {
		for idx, nedge := range nedges {
			if nedge == 1 { // leaf
				next = append(next, idx)
			}
		}
		for _, node := range next {
			removed[node] = true
			nedges[node]--
			for _, near := range adj[node] {
				nedges[near]--
			}
		}
		n -= len(next)
		next = next[:0]
	}

	res := make([]int, 0)
	for idx, didRemove := range removed {
		if !didRemove {
			res = append(res, idx)
		}
	}

	return res
}

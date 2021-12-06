package p1971findifpathexistsingraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validPath(t *testing.T) {
	for _, tc := range []struct {
		n          int
		edges      [][]int
		start, end int
		want       bool
	}{
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
		{6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, validPath(tc.n, tc.edges, tc.start, tc.end))
		})
	}
}

func validPath(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}
	adj := make([][]uint32, n)
	for _, edge := range edges {
		a, b := uint32(edge[0]), uint32(edge[1])
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	seen := make([]bool, n)
	seen[start] = true
	cur := make([]uint32, 1, 1000)
	cur[0] = uint32(start)
	next := make([]uint32, 0, 1000)
	for len(cur) > 0 {
		next = next[:0]
		for _, node := range cur {
			for _, near := range adj[node] {
				if near == uint32(end) {
					return true
				}
				if seen[near] {
					continue
				}
				seen[near] = true
				next = append(next, near)
			}
		}
		cur, next = next, cur
	}
	return false
}

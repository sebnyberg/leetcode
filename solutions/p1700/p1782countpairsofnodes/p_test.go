package p1782countpairsofnodes

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPairs(t *testing.T) {
	for _, tc := range []struct {
		n       int
		edges   [][]int
		queries []int
		want    []int
	}{
		{4, [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {3, 1}}, []int{2, 3}, []int{6, 5}},
		{5, [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}, []int{1, 2, 3, 4, 5}, []int{10, 10, 9, 8, 6}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countPairs(tc.n, tc.edges, tc.queries))
		})
	}
}

func countPairs(n int, edges [][]int, queries []int) []int {
	sharedEdges := make(map[int]map[int]int)
	indeg := make([]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		indeg[a]++
		indeg[b]++
		if a < b {
			a, b = b, a
		}
		if _, exists := sharedEdges[a]; !exists {
			sharedEdges[a] = make(map[int]int)
		}
		sharedEdges[a][b]++
	}

	sortedIndeg := make([]int, n+1)
	copy(sortedIndeg, indeg)
	sort.Ints(sortedIndeg)
	res := make([]int, len(queries))
	for k, q := range queries {
		// Disregarding shared edges, how many combinations of nodes
		// have a total edge count above q?
		for l, r := 1, n; l < r; {
			if sortedIndeg[l]+sortedIndeg[r] > q {
				// Add all nodes pairs from i to j
				// Note that there are j-i+1 nodes, but j-i pairs of nodes containing j
				res[k] += r - l
				r-- // decrease largest count
			} else {
				l++ // increase lowest count (increasing total count)
			}
		}
		for a := 1; a <= n; a++ {
			if _, exists := sharedEdges[a]; !exists {
				continue
			}
			for b, count := range sharedEdges[a] {
				if indeg[a]+indeg[b] > q && indeg[a]+indeg[b]-count <= q {
					res[k]--
				}
			}
		}
	}

	return res
}

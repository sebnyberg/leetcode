package p1489findcriticalpseudocriticaledgesinminspantree

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findCriticalAndPseudoCriticalEdges(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  [][]int
	}{
		{4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}, [][]int{{}, {0, 1, 2, 3}}},
		{5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}, [][]int{{0, 1}, {2, 3, 4, 5}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findCriticalAndPseudoCriticalEdges(tc.n, tc.edges))
		})
	}
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// Positional index is needed in the return values, add it before sort
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// Find MST using Kruskal
	minWeight := mstWeight(n, edges, -1, -1)

	critical := make([]int, 0)
	pseudoCritical := make([]int, 0)

	// For each edge
	for i := range edges {

		// Check the weight with this edge excluded
		wgt := mstWeight(n, edges, i, -1)

		// If weight is -1 (MST not possible to form) or weight increased
		// the edge must be included in the MST (it is critical)
		if wgt == -1 || wgt > minWeight {
			critical = append(critical, edges[i][3])
			continue
		}

		// If weight is == minWeight, we are not sure if this edge was excluded
		// or included. Try force-inclusion to see if weight stays the same
		wgt = mstWeight(n, edges, -1, i)
		if wgt == minWeight {
			pseudoCritical = append(pseudoCritical, edges[i][3])
		}
	}

	// Sort and return
	sort.Ints(critical)
	sort.Ints(pseudoCritical)
	return [][]int{critical, pseudoCritical}
}

func mstWeight(n int, edges [][]int, skipIdx int, forceIncludeIndex int) int {
	dsu := NewDSU(n)

	var weight int
	if forceIncludeIndex != -1 {
		e := edges[forceIncludeIndex]
		dsu.union(e[0], e[1])
		weight += e[2]
	}

	for i, edge := range edges {
		if i == skipIdx {
			continue
		}
		a, b := edge[0], edge[1]
		ra, rb := dsu.find(a), dsu.find(b)
		if ra == rb { // skip when already part of the same component
			continue
		}
		dsu.union(a, b)
		weight += edge[2]
	}

	// Validate that all edges belong to the same component
	a := dsu.find(0)
	for i := range dsu.parent[1:] {
		if dsu.find(1+i) != a {
			return -1
		}
	}
	return weight
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}

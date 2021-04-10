package p1697checkingexistenceofedgelengthlimitedpaths

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distanceLimitedPathsExist(t *testing.T) {
	for _, tc := range []struct {
		n        int
		edgeList [][]int
		queries  [][]int
		want     []bool
	}{
		{5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}, [][]int{{0, 4, 14}, {1, 4, 13}}, []bool{true, false}},
		{3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}}, []bool{false, true}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, distanceLimitedPathsExist(tc.n, tc.edgeList, tc.queries))
		})
	}
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// Append positions to edgeList / queries
	// for i := range edgeList {
	// 	edgeList[i] = append(edgeList[i], i)
	// }
	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	// Sort edgeList and queries by distance
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	// For each query
	res := make([]bool, len(queries))
	var i int
	dsu := NewDSU(n)
	nedges := len(edgeList)
	for _, query := range queries {
		// Connect all edges smaller than the limit
		p, q, limit := query[0], query[1], query[2]

		for ; i < nedges && edgeList[i][2] < limit; i++ {
			// add edge to DSU
			e := edgeList[i]
			dsu.union(e[0], e[1])
		}

		res[query[3]] = dsu.find(p) == dsu.find(q)
	}

	return res
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

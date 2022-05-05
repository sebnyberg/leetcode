package p0685redundantconnection2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	for _, tc := range []struct {
		edges [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}, []int{4, 1}},
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, findRedundantDirectedConnection(tc.edges))
		})
	}
}

func findRedundantDirectedConnection(edges [][]int) []int {
	parentEdgeIndices := make([]int, len(edges)+1)
	for i := range parentEdgeIndices {
		parentEdgeIndices[i] = -1
	}
	for i, edge := range edges {
		if parentEdgeIndices[edge[1]] != -1 { // two parents
			// If removing the current edge leads to no redundant connections
			// return the current index
			if res := findRedundantConnection(edges, i); len(res) == 0 {
				return edges[i]
			}
			// Otherwise return the other candidate
			return edges[parentEdgeIndices[edge[1]]]
		}
		parentEdgeIndices[edge[1]] = i
	}

	// No double parent was found, simply find the redundant connection
	// using the old solution (the redundant connection is the first cycle)
	res := findRedundantConnection(edges, -1)
	return res
}

func findRedundantConnection(edges [][]int, skipIndex int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1)
	var redundant []int
	for i, edge := range edges {
		if i == skipIndex {
			continue
		}
		if dsu.find(edge[0]) == dsu.find(edge[1]) {
			redundant = edge
		}
		dsu.union(edge[0], edge[1])
	}
	return redundant
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

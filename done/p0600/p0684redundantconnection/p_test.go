package p0684redundantconnection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRedundantConnection(t *testing.T) {
	for _, tc := range []struct {
		edges [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, findRedundantConnection(tc.edges))
		})
	}
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1)
	var redundant []int
	for _, edge := range edges {
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

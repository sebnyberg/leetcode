package p0323numberofconnectedcompingraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countComponents(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  int
	}{
		{5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countComponents(tc.n, tc.edges))
		})
	}
}

func countComponents(n int, edges [][]int) int {
	dsu := NewDSU(n)
	for _, edge := range edges {
		dsu.union(edge[0], edge[1])
	}
	seen := make(map[int]struct{})
	for i := 0; i < n; i++ {
		seen[dsu.find(i)] = struct{}{}
	}
	return len(seen)
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

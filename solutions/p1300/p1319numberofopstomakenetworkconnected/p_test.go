package p1319numberofopstomakenetworkconnected

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeConnected(t *testing.T) {
	for _, tc := range []struct {
		n           int
		connections [][]int
		want        int
	}{
		{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}, 0},
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.n, tc.connections), func(t *testing.T) {
			require.Equal(t, tc.want, makeConnected(tc.n, tc.connections))
		})
	}
}

func makeConnected(n int, connections [][]int) int {
	dsu := NewDSU(n)
	navailable := 0
	for _, conn := range connections {
		if dsu.find(conn[0]) == dsu.find(conn[1]) {
			navailable++
		} else {
			dsu.union(conn[0], conn[1])
		}
	}
	ngroups := 0
	seen := make(map[int]struct{})
	for i := 0; i < n; i++ {
		root := dsu.find(i)
		if _, exists := seen[root]; !exists {
			ngroups++
		}
		seen[root] = struct{}{}
	}

	if navailable >= ngroups-1 {
		return ngroups - 1
	}
	return -1
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

func (d *DSU) sameSet(a, b int) bool {
	return d.find(a) == d.find(b)
}

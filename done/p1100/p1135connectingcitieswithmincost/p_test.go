package p1135connectingcitieswithmincost

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCost(t *testing.T) {
	for _, tc := range []struct {
		N           int
		connections [][]int
		want        int
	}{
		{3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}, 6},
		{4, [][]int{{1, 2, 3}, {3, 4, 4}}, -1},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.N, tc.connections), func(t *testing.T) {
			require.Equal(t, tc.want, minimumCost(tc.N, tc.connections))
		})
	}
}

func minimumCost(N int, connections [][]int) int {
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	dsu := NewDSU(N + 1)
	cost := 0
	nconnections := 0
	for _, conn := range connections {
		if dsu.find(conn[0]) == dsu.find(conn[1]) { // already connected
			continue
		}
		dsu.union(conn[0], conn[1])
		cost += conn[2]
		nconnections++
	}
	if nconnections == N-1 {
		return cost
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

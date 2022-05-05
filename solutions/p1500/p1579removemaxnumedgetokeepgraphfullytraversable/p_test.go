package p1579removemaxnumedgetokeepgraphfullytraversable

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumEdgesToRemove(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  int
	}{
		{2, [][]int{{1, 1, 2}, {2, 1, 2}, {3, 1, 2}}, 2},
		{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}, 2},
		{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}, 0},
		{4, [][]int{{3, 1, 2}, {1, 1, 2}, {2, 3, 4}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumEdgesToRemove(tc.n, tc.edges))
		})
	}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	var ntoremove int
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})
	alice := NewDSU(n + 1)
	bob := NewDSU(n + 1)
	for _, edge := range edges {
		t, a, b := edge[0], edge[1], edge[2]
		switch t {
		case 1:
			if alice.find(a) == alice.find(b) {
				ntoremove++
				continue
			}
			alice.union(a, b)
		case 2:
			if bob.find(a) == bob.find(b) {
				ntoremove++
				continue
			}
			bob.union(a, b)
		case 3:
			if bob.find(a) == bob.find(b) || alice.find(a) == alice.find(b) {
				ntoremove++
				continue
			}
			alice.union(a, b)
			bob.union(a, b)
		}
	}
	if alice.size[alice.find(1)] != n || bob.size[bob.find(1)] != n {
		return -1
	}
	return ntoremove
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

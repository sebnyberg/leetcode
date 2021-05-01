package daily

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
	d := NewDSU(n)
	for _, edge := range edges {
		d.Union(edge[0], edge[1])
	}
	components := make(map[int]struct{})
	for i := 0; i < n; i++ {
		components[d.Find(i)] = struct{}{}
	}
	return len(components)
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
	}
	return dsu
}

func (d *DSU) Find(a int) int {
	if d.parent[a] == a {
		return a
	}
	r := d.Find(d.parent[a])
	d.parent[a] = r
	return r
}

func (d *DSU) Union(a, b int) {
	ra, rb := d.Find(a), d.Find(b)
	d.parent[ra] = rb
}

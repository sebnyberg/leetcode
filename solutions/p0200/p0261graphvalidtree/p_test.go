package p0261graphvalidtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validTree(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  bool
	}{
		{5, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 4}}, false},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},
		{5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, validTree(tc.n, tc.edges))
		})
	}
}

func validTree(n int, edges [][]int) bool {
	// A valid tree is not part of the DSU before being added
	dsu := DSU{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		ra, rb := dsu.find(a), dsu.find(b)
		if ra == rb {
			return false
		}
		dsu.parent[ra] = rb
	}

	// Now all items should be in the same set
	first := dsu.find(0)
	for i := 1; i < len(dsu.parent); i++ {
		if dsu.find(i) != first {
			return false
		}
	}

	return true
}

type DSU struct {
	parent []int
}

func (d *DSU) find(a int) int {
	if d.parent[a] != a {
		return d.find(d.parent[a])
	}
	return a
}

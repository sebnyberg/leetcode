package p0959regionscutbyslashes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_regionsBySlashes(t *testing.T) {
	for _, tc := range []struct {
		grid []string
		want int
	}{
		{[]string{
			"//",
			"/ ",
		}, 3},
		{[]string{
			" /",
			"/ ",
		}, 2},
		{[]string{
			" /",
			"  ",
		}, 1},
		{[]string{
			"\\/",
			"/\\",
		}, 4},
		{[]string{
			"/\\",
			"\\/",
		}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, regionsBySlashes(tc.grid))
		})
	}
}

func regionsBySlashes(grid []string) int {
	dsu := NewDSU()
	m, n := len(grid), len(grid[0])

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dsu.parent[position{i, j}] = position{i, j}
		}
	}

	// Edges form a set
	for i := 0; i < m; i++ {
		dsu.union(position{i, 0}, position{i + 1, 0})
		dsu.union(position{i, n}, position{i + 1, n})
	}
	for i := 0; i < n; i++ {
		dsu.union(position{0, i}, position{0, i + 1})
		dsu.union(position{m, i}, position{m, i + 1})
	}

	// For each cell in the grid,
	nregions := 1
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == ' ' {
				continue
			}
			var from, to position
			if ch == '\\' {
				from = position{i, j}
				to = position{i + 1, j + 1}
			} else {
				from = position{i, j + 1}
				to = position{i + 1, j}
			}
			// If both endpoints are in the same set, this connection
			// cuts a region in half
			if dsu.find(from) == dsu.find(to) {
				nregions++
			}
			dsu.union(from, to)
		}
	}

	return nregions
}

type position struct {
	i, j int
}

type DSU struct {
	parent map[position]position
}

func NewDSU() *DSU {
	dsu := &DSU{
		parent: make(map[position]position),
	}
	return dsu
}

func (d *DSU) find(a position) position {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b position) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		d.parent[b] = a
	}
}

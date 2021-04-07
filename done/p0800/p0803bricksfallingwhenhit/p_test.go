package p0803bricksfallingwhenhit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hitBricks(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		hits [][]int
		want []int
	}{
		{[][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}}, []int{0, 3, 0}},
		{[][]int{{1}, {1}, {1}, {1}, {1}}, [][]int{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}}, []int{1, 0, 1, 0, 0}},
		{[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}}, [][]int{{1, 1}, {1, 0}}, []int{0, 0}},
		{[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}, []int{2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, hitBricks(tc.grid, tc.hits))
		})
	}
}

func hitBricks(grid [][]int, hits [][]int) []int {
	// For the parent array, index 0 means outside the board on the top
	// index 1,...,m is the top row of bricks
	// index m+1,...,2m is the second row, and so on
	// Neighbours for a given brick will be +-1 and +- m (unless out of bounds)

	// Instead of considering hits by falling rocks, reverse time so that
	// an "unhit" rock can connect a previously disconnected island.
	// If the connection happens, it means that island would've been removed,
	// this yielding the result

	m, n := len(grid), len(grid[0])
	gridCopy := make([][]int, m)
	for i := range gridCopy {
		gridCopy[i] = make([]int, n)
		copy(gridCopy[i], grid[i])
	}

	// Hit all rocks
	for _, hit := range hits {
		gridCopy[hit[0]][hit[1]] = 0
	}

	dsu := NewDSU(m*n + 1)

	// Create all connections
	for i := 0; i < m; i++ {
		for j := range gridCopy[i] {
			if gridCopy[i][j] == 1 {
				// connect to all nearby nodes
				pos := n*i + j + 1
				if j > 0 && gridCopy[i][j-1] == 1 { // left
					dsu.union(pos, pos-1)
				}
				if j < n-1 && gridCopy[i][j+1] == 1 { // right
					dsu.union(pos, pos+1)
				}
				if i == 0 {
					dsu.union(pos, 0)
				} else if gridCopy[i-1][j] == 1 {
					dsu.union(pos, pos-n)
				}
				if i < m-1 && gridCopy[i+1][j] == 1 { // below
					dsu.union(pos, pos+n)
				}
			}
		}
	}

	// Re-introduce a previously hit stone
	// If a connection is made, put the size of the current group in the response
	res := make([]int, len(hits))
	for t := len(hits) - 1; t >= 0; t-- {
		i, j := hits[t][0], hits[t][1]
		if grid[i][j] == 0 {
			continue
		}
		gridCopy[i][j] = 1
		pos := n*i + j + 1
		// Any group that is found that does not have zero as its root node
		// would be introduced to the top
		sizeBefore := dsu.size[0]
		if j > 0 && gridCopy[i][j-1] == 1 { // left
			dsu.union(pos, pos-1)
		}
		if j < n-1 && gridCopy[i][j+1] == 1 { // right
			dsu.union(pos, pos+1)
		}
		if i == 0 { // above
			dsu.union(pos, 0)
		} else if gridCopy[i-1][j] == 1 {
			dsu.union(pos, pos-n)
		}
		if i < m-1 && gridCopy[i+1][j] == 1 { // below
			dsu.union(pos, pos+n)
		}
		res[t] = max(dsu.size[0]-sizeBefore-1, 0)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		if a == 0 { // always prefer connecting toward the root
			a, b = b, a
		}
		d.parent[a] = b
		d.size[b] += d.size[a]
	}
}

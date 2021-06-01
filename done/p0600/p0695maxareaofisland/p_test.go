package p0695maxareaofisland

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAreaOfIsland(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maxAreaOfIsland(tc.grid))
		})
	}
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	islands := NewDSU(m * n)

	// For each position, add it to the disjoint set union
	seen := make([]bool, m*n)
	var idx int
	for i := range grid {
		for j, cell := range grid[i] {
			if cell == 0 {
				idx++
				continue
			}
			islands.maxSize = max(islands.maxSize, 1)
			// join the cell with any non-zero neighbours
			for _, nei := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j + 1}, {i, j - 1},
			} {
				if nei[0] < 0 || nei[1] < 0 || nei[0] >= m || nei[1] >= n || grid[nei[0]][nei[1]] != 1 {
					continue
				}
				neiIdx := nei[0]*n + nei[1]
				if seen[neiIdx] {
					continue
				}
				islands.union(idx, neiIdx)
			}
			seen[idx] = true
			idx++
		}
	}

	return islands.maxSize
}

type dsu struct {
	parent  []int
	size    []int
	maxSize int
}

func NewDSU(n int) *dsu {
	// Create dsu
	d := &dsu{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range d.parent {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *dsu) find(x int) int {
	if d.parent[x] == x {
		return x
	}
	root := d.find(d.parent[x])
	d.parent[x] = root // path compression
	return root
}

func (d *dsu) union(x, y int) {
	rx := d.find(x)
	ry := d.find(y)
	if rx == ry {
		return
	}
	if d.size[rx] < d.size[ry] {
		rx, ry = ry, rx
	}
	d.parent[ry] = rx
	d.size[rx] += d.size[ry]
	d.maxSize = max(d.maxSize, d.size[rx])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

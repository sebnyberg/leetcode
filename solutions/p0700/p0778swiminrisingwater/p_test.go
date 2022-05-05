package p0778swiminrisingwater

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_swimInWater(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{3, 2}, {0, 1}}, 3},
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, swimInWater(tc.grid))
		})
	}
}

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	cells := make([]gridCell, 0, m*n)
	for i := range grid {
		for j, v := range grid[i] {
			cells = append(cells, gridCell{i, j, v})
		}
	}
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].elevation < cells[j].elevation
	})

	dsu := NewDSU(m * n)
	for _, c := range cells {
		for _, near := range [][2]int{
			{c.i + 1, c.j}, {c.i - 1, c.j}, {c.i, c.j - 1}, {c.i, c.j + 1},
		} {
			i, j := near[0], near[1]
			if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] > c.elevation {
				continue
			}
			dsu.union(i*n+j, c.i*n+c.j)
			if dsu.find(0) == dsu.find(m*n-1) {
				return c.elevation
			}
		}
	}
	return -1
}

type gridCell struct {
	i, j      int
	elevation int
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

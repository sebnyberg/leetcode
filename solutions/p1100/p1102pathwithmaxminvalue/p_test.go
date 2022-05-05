package p1102pathwithmaxminvalue

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumMinimumPath(t *testing.T) {
	for _, tc := range []struct {
		A    [][]int
		want int
	}{
		{[][]int{{5, 4, 5}, {1, 2, 6}, {7, 4, 6}}, 4},
		{[][]int{{2, 2, 1, 2, 2, 2}, {1, 2, 2, 2, 1, 2}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, maximumMinimumPath(tc.A))
		})
	}
}

func maximumMinimumPath(A [][]int) int {
	m, n := len(A), len(A[0])
	cells := make([]gridCell, m*n)
	for i := range A {
		for j, v := range A[i] {
			cells = append(cells, gridCell{i, j, v})
		}
	}
	sort.Slice(cells, func(i, j int) bool { return cells[i].val > cells[j].val })
	dsu := NewDSU(m, n)
	visited := make(map[[2]int]struct{})
	for i, cell := range cells {
		visited[[2]int{cell.i, cell.j}] = struct{}{}
		if i == 0 {
			continue
		}
		for _, near := range getNear(cell.i, cell.j, m, n) {
			if _, exists := visited[[2]int{near[0], near[1]}]; exists {
				dsu.union(cell.i, cell.j, near[0], near[1])
			}
		}
		if dsu.find(0, 0) == dsu.find(m-1, n-1) {
			return cell.val
		}
	}
	return -1
}

type gridCell struct {
	i, j int
	val  int
}

func getNear(i, j, m, n int) [][2]int {
	nearby := make([][2]int, 0, 4)
	for _, near := range [][2]int{
		{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
	} {
		if near[0] < 0 || near[1] < 0 || near[0] >= m || near[1] >= n {
			continue
		}
		nearby = append(nearby, near)
	}
	return nearby
}

type DSU struct {
	parent [][][2]int
	size   [][]int
}

func NewDSU(m int, n int) *DSU {
	dsu := &DSU{
		parent: make([][][2]int, m),
		size:   make([][]int, m),
	}
	for i := 0; i < m; i++ {
		dsu.parent[i] = make([][2]int, n)
		dsu.size[i] = make([]int, n)
		for j := range dsu.parent[i] {
			dsu.parent[i][j] = [2]int{i, j}
			dsu.size[i][j] = 1
		}
	}
	return dsu
}

func (d *DSU) find(i, j int) [2]int {
	if d.parent[i][j] == [2]int{i, j} {
		return [2]int{i, j}
	}
	root := d.find(d.parent[i][j][0], d.parent[i][j][1])
	d.parent[i][j] = root
	return root
}

func (d *DSU) union(ai, aj, bi, bj int) {
	a := d.find(ai, aj)
	b := d.find(bi, bj)
	if a != b {
		if d.size[a[0]][a[1]] < d.size[b[0]][b[1]] {
			a, b = b, a
		}
		d.parent[b[0]][b[1]] = a
		d.size[a[0]][a[1]] += d.size[b[0]][b[1]]
	}
}

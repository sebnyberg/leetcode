package p1905countsubislands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubIslands(t *testing.T) {
	for _, tc := range []struct {
		grid1 [][]int
		grid2 [][]int
		want  int
	}{
		{
			[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
			3,
		},
		{
			[][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
			[][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}},
			2,
		},
		{
			[][]int{{1, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0}, {1, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 0}},
			[][]int{{1, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1}, {1, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1}, {1, 0, 0, 1, 0, 0}},
			0,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid1), func(t *testing.T) {
			require.Equal(t, tc.want, countSubIslands(tc.grid1, tc.grid2))
		})
	}
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	islands := NewDSU(m * n)

	getpos := func(i, j int) int {
		return i*n + j
	}

	// Put first grid into the disjoint set union
	for i := range grid1 {
		for j, cell := range grid1[i] {
			if cell == 0 {
				continue
			}

			// Cell is 1, join with nearby 1s
			for _, near := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j + 1}, {i, j - 1},
			} {
				if near[0] < 0 || near[0] >= m || near[1] < 0 || near[1] >= n || grid1[near[0]][near[1]] == 0 {
					continue
				}
				aa, bb := getpos(i, j), getpos(near[0], near[1])
				islands.union(aa, bb)
			}
		}
	}

	subIslands := NewDSU(m * n)

	// Put second grid into the disjoint set union
	for i := range grid2 {
		for j, cell := range grid2[i] {
			if cell == 0 {
				continue
			}

			// Cell is '1', join with nearby 1s
			for _, near := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j + 1}, {i, j - 1},
			} {
				if near[0] < 0 || near[0] >= m || near[1] < 0 || near[1] >= n || grid2[near[0]][near[1]] == 0 {
					continue
				}
				subIslands.union(getpos(i, j), getpos(near[0], near[1]))
			}
		}
	}

	// Finally, for each unique island in sub-islands, ensure that all cells
	// within that island map to the same island in the islands DSU
	subIslandCells := make(map[int][]int)
	for idx, rootIdx := range subIslands.parent {
		if grid2[idx/n][idx%n] == 1 {
			ri := subIslands.find(rootIdx)
			subIslandCells[ri] = append(subIslandCells[ri], idx)
		}
	}

	var res int
	for _, subIslandCellIndices := range subIslandCells {
		firstIdx := subIslandCellIndices[0]
		if grid1[firstIdx/n][firstIdx%n] != 1 {
			continue
		}
		firstIsland := islands.find(subIslandCellIndices[0])
		for _, idx := range subIslandCellIndices[1:] {
			ri := islands.find(idx)
			if ri != firstIsland || grid1[idx/n][idx%n] != 1 {
				goto ContinueLoop
			}
		}
		res++
	ContinueLoop:
	}
	return res
}

type dsu struct {
	parent []int
}

func NewDSU(n int) *dsu {
	d := &dsu{
		parent: make([]int, n),
	}
	for i := range d.parent {
		d.parent[i] = i
	}
	return d
}

func (d *dsu) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *dsu) union(a, b int) {
	ra := d.find(a)
	rb := d.find(b)
	if ra != rb {
		d.parent[ra] = rb
	}
}

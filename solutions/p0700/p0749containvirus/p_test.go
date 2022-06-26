package p4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containVirus(t *testing.T) {
	for _, tc := range []struct {
		isInfected [][]int
		want       int
	}{
		{
			[][]int{
				{0, 1, 0, 1, 1, 1, 1, 1, 1, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 0, 0, 0, 1, 0},
				{0, 0, 0, 1, 1, 0, 0, 1, 1, 0},
				{0, 1, 0, 0, 1, 0, 1, 1, 0, 1},
				{0, 0, 0, 1, 0, 1, 0, 1, 1, 1},
				{0, 1, 0, 0, 1, 0, 0, 1, 1, 0},
				{0, 1, 0, 1, 0, 0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0, 1, 1, 0, 0, 1},
				{1, 0, 1, 1, 0, 1, 0, 1, 0, 1},
			},
			38,
		},
		{
			[][]int{
				{0, 1, 0, 0, 0, 0, 0, 1},
				{0, 1, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			10,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			4,
		},
		{
			[][]int{
				{1, 1, 1, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 0, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 0, 0, 0, 0, 0},
			},
			13,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.isInfected), func(t *testing.T) {
			require.Equal(t, tc.want, containVirus(tc.isInfected))
		})
	}
}

const (
	empty     = 0
	infected  = 1
	contained = 2
)

func containVirus(isInfected [][]int) int {
	res := progress(isInfected)
	return res
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func spread(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != contained
	}
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	// Find all infected cells
	infectedCells := [][]int{}
	var k int
	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seen[i][j] || grid[i][j] == contained || grid[i][j] == empty {
				continue
			}
			infectedCells = append(infectedCells, []int{i, j})
			seen[i][j] = true
			for ; k < len(infectedCells); k++ {
				for _, d := range dirs {
					ii := infectedCells[k][0] + d[0]
					jj := infectedCells[k][1] + d[1]
					if !ok(ii, jj) || seen[ii][jj] || grid[ii][jj] != infected {
						continue
					}
					seen[ii][jj] = true
					infectedCells = append(infectedCells, []int{ii, jj})
				}
			}
		}
	}

	// Infect all neighbours of infected cells that are not contained.
	for _, cell := range infectedCells {
		for _, d := range dirs {
			ii := cell[0] + d[0]
			jj := cell[1] + d[1]
			if !ok(ii, jj) || grid[ii][jj] == contained || grid[ii][jj] == infected {
				continue
			}
			grid[ii][jj] = infected
			count++
		}
	}
	return count
}

func progress(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != contained
	}
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	regionCells := make([][2]int, 0, 10)

	// Find A cell which is part of a region that threatens the most cells
	var maxThreatened int
	var maxThreatenedCells [][2]int
	for i := range grid {
		for j := range grid[i] {
			if !ok(i, j) || seen[i][j] || grid[i][j] != infected {
				continue
			}
			// Create a copy of the grid
			cpy := make([][]int, m)
			for i := range cpy {
				cpy[i] = make([]int, n)
				copy(cpy[i], grid[i])
			}
			// Collect infected cells
			regionCells = regionCells[:0]
			regionCells = append(regionCells, [2]int{i, j})
			seen[i][j] = true
			cpy[i][j] = contained
			for k := 0; k < len(regionCells); k++ {
				for _, d := range dirs {
					ii := regionCells[k][0] + d[0]
					jj := regionCells[k][1] + d[1]
					if !ok(ii, jj) || cpy[ii][jj] == contained || cpy[ii][jj] == empty {
						continue
					}
					seen[ii][jj] = true
					cpy[ii][jj] = contained
					regionCells = append(regionCells, [2]int{ii, jj})
				}
			}
			// Simulate spread
			var count int
			for _, cell := range regionCells {
				for _, d := range dirs {
					ii := cell[0] + d[0]
					jj := cell[1] + d[1]
					if !ok(ii, jj) || cpy[ii][jj] == contained || cpy[ii][jj] == infected {
						continue
					}
					cpy[ii][jj] = infected
					count++
				}
			}
			if count > maxThreatened {
				maxThreatened = count
				maxThreatenedCells = make([][2]int, len(regionCells))
				copy(maxThreatenedCells, regionCells)
			}
		}
	}

	if maxThreatened == 0 {
		return 0
	}

	// Count walls
	res := countWalls(grid, maxThreatenedCells)

	// Mark maxThreatenedCells as contained
	for _, cell := range maxThreatenedCells {
		grid[cell[0]][cell[1]] = contained
	}

	// Spread
	if spread(grid) == 0 {
		return res
	}

	return res + progress(grid)
}

func countWalls(grid [][]int, cells [][2]int) int {
	m, n := len(grid), len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != contained
	}
	// Count walls around cells
	var wallCount int
	for _, cell := range cells {
		for _, d := range dirs {
			ii, jj := cell[0]+d[0], cell[1]+d[1]
			if !ok(ii, jj) || grid[ii][jj] != empty {
				continue
			}
			wallCount++
		}
	}
	return wallCount
}

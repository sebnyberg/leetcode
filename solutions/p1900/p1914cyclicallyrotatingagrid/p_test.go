package p1914cyclicallyrotatingagrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotateGrid(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		k    int
		want [][]int
	}{
		{[][]int{{40, 10}, {30, 20}}, 1, [][]int{{10, 20}, {40, 30}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 2, [][]int{{3, 4, 8, 12}, {2, 11, 10, 16}, {1, 7, 6, 15}, {5, 9, 13, 14}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, rotateGrid(tc.grid, tc.k))
		})
	}
}

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	nextGrid := make([][]int, m)
	for i := range nextGrid {
		nextGrid[i] = make([]int, n)
		copy(nextGrid[i], grid[i])
	}
	curGrid := make([][]int, m)
	for i := range grid {
		curGrid[i] = make([]int, n)
		copy(curGrid[i], grid[i])
	}
	for offset := 0; offset < min(m/2, n/2); offset++ {
		mod := (m - (offset * 2)) * 2
		mod += (n - (offset * 2) - 2) * 2
		modK := k % mod
		for i := 0; i < modK; i++ {
			for l := offset; l < n-1-offset; l++ {
				nextGrid[offset][l] = curGrid[offset][l+1]
			}
			for l := offset; l < m-1-offset; l++ {
				nextGrid[l][n-1-offset] = curGrid[l+1][n-1-offset]
			}
			for l := n - 1 - offset; l > offset; l-- {
				nextGrid[m-1-offset][l] = curGrid[m-1-offset][l-1]
			}
			for l := m - 1 - offset; l > offset; l-- {
				nextGrid[l][offset] = curGrid[l-1][offset]
			}
			curGrid, nextGrid = nextGrid, curGrid
		}
		// Copy columns
		for i := offset; i < m-offset; i++ {
			grid[i][offset] = curGrid[i][offset]
			grid[i][n-1-offset] = curGrid[i][n-1-offset]
		}
		// Copy rows
		for i := offset; i < n-offset; i++ {
			grid[offset][i] = curGrid[offset][i]
			grid[m-1-offset][i] = curGrid[m-1-offset][i]
		}
	}
	return grid
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

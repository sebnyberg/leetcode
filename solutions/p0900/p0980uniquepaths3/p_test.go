package p0980uniquepaths3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePathsIII(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}, 4},
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, uniquePathsIII(tc.grid))
		})
	}
}

func uniquePathsIII(grid [][]int) (npaths int) {
	var pathlen int
	var startI, startJ int
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 0:
				pathlen++
			case 1:
				startI, startJ = i, j
			}
		}
	}

	grid[startI][startJ] = 0
	return findUnique(startI, startJ, 0, pathlen+1, grid)
}

func findUnique(i int, j int, cursteps int, wantsteps int, grid [][]int) (paths int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	switch grid[i][j] {
	case 2:
		if cursteps == wantsteps {
			return 1
		}
		return 0
	case -1:
		return 0
	case 1:
		return 0
	}
	grid[i][j] = 1
	cursteps++
	paths += findUnique(i-1, j, cursteps, wantsteps, grid)
	paths += findUnique(i+1, j, cursteps, wantsteps, grid)
	paths += findUnique(i, j-1, cursteps, wantsteps, grid)
	paths += findUnique(i, j+1, cursteps, wantsteps, grid)
	cursteps--
	grid[i][j] = 0
	return paths
}

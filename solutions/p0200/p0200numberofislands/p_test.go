package p0200numberofislands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numIslands(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, numIslands(tc.grid))
		})
	}
}

func clearIsland(grid [][]byte, n, m, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	clearIsland(grid, n, m, i+1, j)
	clearIsland(grid, n, m, i-1, j)
	clearIsland(grid, n, m, i, j+1)
	clearIsland(grid, n, m, i, j-1)
}

func numIslands(grid [][]byte) int {
	n, m := len(grid[0]), len(grid)
	nislands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			nislands++
			clearIsland(grid, n, m, i, j)
		}
	}
	return nislands
}

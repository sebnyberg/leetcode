package p0062uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePaths(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, uniquePaths(tc.m, tc.n))
		})
	}
}

func uniquePaths(m int, n int) int {
	// DP-solution
	// Since the robot can only move down or right, the number of
	// ways to reach the top and left cells in the grid is one
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
		grid[i][0] = 1
	}
	for i := range grid[0] {
		grid[0][i] = 1
	}

	// The number of ways the robot can reach any other point in the
	// grid is the sum of the ways it can reach the position above,
	// and to the left
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

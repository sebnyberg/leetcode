package p0063uniquepaths2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePaths(t *testing.T) {
	for _, tc := range []struct {
		obstacleGrid [][]int
		want         int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.obstacleGrid), func(t *testing.T) {
			require.Equal(t, tc.want, uniquePathsWithObstacles(tc.obstacleGrid))
		})
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// DP-solution
	// Since the robot can only move down or right, the number of
	// ways to reach the top and left cells in the grid is one
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for i := range obstacleGrid {
		if obstacleGrid[i][0] == 1 {
			break
		}
		grid[i][0] = 1
	}
	for i := range obstacleGrid[0] {
		if obstacleGrid[0][i] == 1 {
			break
		}
		grid[0][i] = 1
	}

	// The number of ways the robot can reach any other point in the
	// grid is the sum of the ways it can reach the position above,
	// and to the left
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

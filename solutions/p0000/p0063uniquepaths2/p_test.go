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
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

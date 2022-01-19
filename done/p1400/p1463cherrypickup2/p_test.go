package p1463cherrypickup2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_cherryPickup(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}, 24},
		{[][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}, 28},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, cherryPickup(tc.grid))
		})
	}
}

func cherryPickup(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt32
			}
		}
	}
	ok := func(pos int) bool {
		return pos >= 0 && pos < m
	}

	dp[0][0][m-1] = grid[0][0] + grid[0][m-1]
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				maxRes := math.MinInt32
				for l := -1; l <= 1; l++ {
					for ll := -1; ll <= 1; ll++ {
						if !ok(j+l) || !ok(k+ll) {
							continue
						}
						res := grid[i][j] + grid[i][k]
						if j == k {
							res /= 2
						}
						maxRes = max(maxRes, res+dp[i-1][j+l][k+ll])
					}
				}
				dp[i][j][k] = maxRes
			}
		}
	}
	maxRes := math.MinInt32
	for j := 0; j < m; j++ {
		for k := 0; k < m; k++ {
			maxRes = max(maxRes, dp[n-1][j][k])
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

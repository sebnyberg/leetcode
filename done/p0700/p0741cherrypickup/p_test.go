package p0741cherrypickup

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
		{[][]int{
			{0, 1, -1},
			{1, 0, -1},
			{1, 1, 1},
		}, 5,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, cherryPickup(tc.grid))
		})
	}
}

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt32
			}
		}
	}
	return max(0, helper(dp, grid, m, n, 0, 0, 0))
}

func helper(dp [][][]int, grid [][]int, m, n, r1, c1, c2 int) int {
	r2 := r1 + c1 - c2
	switch {
	case
		r1 >= m, c1 >= n, c2 >= n, r2 >= m,
		grid[r1][c1] == -1,
		grid[r2][c2] == -1:
		return math.MinInt32
	case r1 == n-1 && c1 == m-1:
		return grid[r1][c1]
	case dp[r1][c1][c2] != math.MinInt32:
		return dp[r1][c1][c2]
	}
	cherries := grid[r1][c1]
	if c1 != c2 {
		cherries += grid[r2][c2]
	}
	cherries += max(
		max(helper(dp, grid, m, n, r1+1, c1, c2), helper(dp, grid, m, n, r1, c1+1, c2)),
		max(helper(dp, grid, m, n, r1+1, c1, c2+1), helper(dp, grid, m, n, r1, c1+1, c2+1)),
	)
	dp[r1][c1][c2] = cherries
	return cherries
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

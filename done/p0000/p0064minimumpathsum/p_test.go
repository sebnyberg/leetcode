package p0064minimumpathsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPathSum(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minPathSum(tc.grid))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
		if i == 0 {
			sum[i][0] = grid[i][0]
			continue
		}
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}
	return sum[m-1][n-1]
}

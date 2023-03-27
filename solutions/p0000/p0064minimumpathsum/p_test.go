package p0064minimumpathsum

import (
	"fmt"
	"math"
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

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	prev := make([]int, n+1)
	curr := make([]int, n+1)
	for j, v := range grid[0] {
		prev[j+1] = prev[j] + v
	}
	for i := 1; i < m; i++ {
		curr[0] = math.MaxInt32
		for j := 0; j < n; j++ {
			curr[j+1] = grid[i][j] + min(curr[j], prev[j+1])
		}
		prev, curr = curr, prev
	}
	return prev[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

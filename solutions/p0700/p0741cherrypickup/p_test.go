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
		{
			[][]int{
				{1, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 1},
				{1, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 1},
			},
			15,
		},
		{
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{-1, -1, 1, 1, -1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, -1, 1},
			},
			14,
		},
		{
			[][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1},
			},
			5,
		},
		{
			[][]int{
				{1, 1, -1},
				{1, -1, 1},
				{-1, 1, 1},
			},
			0,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, cherryPickup(tc.grid))
		})
	}
}

const (
	wall   = -1
	empty  = 0
	cherry = 1
)

func cherryPickup(grid [][]int) int {
	n := len(grid)

	dp := make([][][]int16, n)
	for i1 := range dp {
		dp[i1] = make([][]int16, n)
		for j1 := range dp[i1] {
			dp[i1][j1] = make([]int16, n)
			for i2 := range dp[i1][j1] {
				dp[i1][j1][i2] = -1
			}
		}
	}

	res := visit(dp, grid, 0, 0, 0, n)
	return max(0, res)
}

func visit(dp [][][]int16, grid [][]int, i1, j1, i2, n int) int {
	// Edge-cases
	j2 := i1 + j1 - i2
	if i1 < 0 || j1 < 0 || i2 < 0 || j2 < 0 ||
		i1 >= n || j1 >= n || i2 >= n || j2 >= n ||
		grid[i1][j1] == wall || grid[i2][j2] == wall {
		return math.MinInt16
	}
	if dp[i1][j1][i2] != -1 {
		return int(dp[i1][j1][i2])
	}

	// Account for cherries in current positions (if any)
	cherries := grid[i1][j1] + grid[i2][j2]
	if i1 == i2 && j1 == j2 {
		cherries -= grid[i1][j1]
	}

	res := math.MinInt16
	if i1 != n-1 || j1 != n-1 {
		res = max(res, visit(dp, grid, i1+1, j1, i2+1, n))
		res = max(res, visit(dp, grid, i1, j1+1, i2+1, n))
		res = max(res, visit(dp, grid, i1+1, j1, i2, n))
		res = max(res, visit(dp, grid, i1, j1+1, i2, n))
	} else {
		res = 0
	}

	// Try all possible paths
	if res != math.MinInt16 {
		res += cherries
	}
	dp[i1][j1][i2] = int16(res)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1219pathwithmaximumgold

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_getMaximumGold(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,6,0],[5,8,7],[0,9,0]]"),
			24,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getMaximumGold(tc.grid))
		})
	}
}

func getMaximumGold(grid [][]int) int {
	// Exhaustively search and memoize best results for each given state.
	//
	// O(n^4)
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			v := grid[i][j]
			grid[i][j] = 0
			res = max(res, v+visit(grid, i, j))
			grid[i][j] = v
		}
	}
	return res
}

var dirs = [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

func visit(grid [][]int, i, j int) int {
	m := len(grid)
	n := len(grid[0])
	var res int
	for _, d := range dirs {
		ii := i + d[0]
		jj := j + d[1]
		if ii < 0 || jj < 0 || ii >= m || jj >= n || grid[ii][jj] == 0 {
			continue
		}
		v := grid[ii][jj]
		grid[ii][jj] = 0
		res = max(res, v+visit(grid, ii, jj))
		grid[ii][jj] = v
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0807maxincreasetokeepcityskyline

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]"), 35},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maxIncreaseKeepingSkyline(tc.grid))
		})
	}
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxRows := make([]int, m)
	maxCols := make([]int, n)
	for i := range grid {
		for j := range grid[i] {
			maxRows[i] = max(maxRows[i], grid[i][j])
			maxCols[j] = max(maxCols[j], grid[i][j])
		}
	}
	var sum int
	for i := range grid {
		for j := range grid[i] {
			sum += min(maxRows[i], maxCols[j]) - grid[i][j]
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

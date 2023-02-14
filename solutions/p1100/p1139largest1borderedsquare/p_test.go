package p1139largest1borderedsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largest1BorderedSquare(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 9},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, largest1BorderedSquare(tc.grid))
		})
	}
}

func largest1BorderedSquare(grid [][]int) int {
	// There aren't that many possible squares in the grid.
	//
	// Given a top-left position of (i,j), there are only min(m-i,n-j) possible
	// squares.
	//
	// So, we can probably just try em all.
	//
	m := len(grid)
	n := len(grid[0])

	// Capture prefix sums
	rows := make([][]int, m)
	for i := range rows {
		rows[i] = make([]int, n+1)
	}
	cols := make([][]int, n)
	for i := range cols {
		cols[i] = make([]int, m+1)
	}
	for i := range grid {
		for j := range grid[i] {
			rows[i][j+1] = rows[i][j] + grid[i][j]
			cols[j][i+1] = cols[j][i] + grid[i][j]
		}
	}

	// For each position, try all sizes of squares
	var res int
	for i := range grid {
		for j := range grid[i] {
			for k := 1; i+k <= m && j+k <= n; k++ {
				// Check cols/rows
				if rows[i][j+k]-rows[i][j] != k {
					continue
				}
				if rows[i+k-1][j+k]-rows[i+k-1][j] != k {
					continue
				}
				if cols[j][i+k]-cols[j][i] != k {
					continue
				}
				if cols[j+k-1][i+k]-cols[j+k-1][i] != k {
					continue
				}
				res = max(res, k*k)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

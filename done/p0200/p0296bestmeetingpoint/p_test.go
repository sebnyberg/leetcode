package p0296bestmeetingpoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minTotalDistance(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}, 6},
		{[][]int{{1, 1}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minTotalDistance(tc.grid))
		})
	}
}

func minTotalDistance(grid [][]int) int {
	// The median minimizes the manhattan distance for a 1d-field
	m, n := len(grid), len(grid[0])

	// Rows
	rowVals := make([]int, 0, m)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				rowVals = append(rowVals, i)
			}
		}
	}

	// Cols
	colVals := make([]int, 0, n)
	for i := range grid[0] {
		for j := range grid {
			if grid[j][i] == 1 {
				colVals = append(colVals, i)
			}
		}
	}

	midRow := rowVals[len(rowVals)/2]
	midCol := colVals[len(colVals)/2]

	// Calculate distance to row
	d := 0
	for _, row := range rowVals {
		d += abs(row - midRow)
	}

	for _, col := range colVals {
		d += abs(col - midCol)
	}
	return d
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package p0688knightprobabilityinchessboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_knightProbability(t *testing.T) {
	for _, tc := range []struct {
		n, k, row, column int
		want              float64
	}{
		{3, 1, 1, 2, 0.25},
		{3, 2, 0, 0, 0.0625},
		{1, 0, 0, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, knightProbability(tc.n, tc.k, tc.row, tc.column))
		})
	}
}

func knightProbability(n int, k int, row int, column int) float64 {
	// The trick is to visit each cell on the board, not use e.g. dfs
	var board [25][25]float64
	board[row][column] = 1
	dirs := [][]int{
		{-2, 1}, {-1, 2}, {2, 1}, {1, 2}, {-2, -1}, {-1, -2}, {2, -1}, {1, -2},
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n
	}
	var outside float64
	for i := 0; i < k; i++ {
		var next [25][25]float64
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				for _, d := range dirs {
					rr, cc := r+d[0], c+d[1]
					if !ok(rr, cc) {
						outside += board[r][c] * 1 / 8
					} else {
						next[rr][cc] += board[r][c] * 1 / 8
					}
				}
			}
		}
		board = next
	}
	res := 1 - outside
	return res
}

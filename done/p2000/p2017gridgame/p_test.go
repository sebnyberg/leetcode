package p2017gridgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gridGame(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int64
	}{
		{[][]int{{2, 5, 4}, {1, 5, 1}}, 4},
		{[][]int{{3, 3, 1}, {8, 5, 2}}, 4},
		{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, gridGame(tc.grid))
		})
	}
}

func gridGame(grid [][]int) int64 {
	// The first robot has the ability to create two segments of length zero to
	// n-1, one in the top right corner, and one in the bottom left.
	// The goal is to minimize the value of that segment.
	n := len(grid[0])
	var bottomSum, topSum int
	for i := 0; i < n-1; i++ {
		bottomSum += grid[1][i]
	}
	minSum := bottomSum
	for i := n - 1; i > 0; i-- {
		// Add top right cell to topsum and remove from the bottom
		topSum += grid[0][i]
		bottomSum -= grid[1][i-1]
		minSum = min(minSum, max(bottomSum, topSum))
	}
	return int64(minSum)
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

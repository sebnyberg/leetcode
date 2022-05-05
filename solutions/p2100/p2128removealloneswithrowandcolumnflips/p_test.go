package p2128removealloneswithrowandcolumnflips

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeOnes(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want bool
	}{
		{[][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, true},
		{[][]int{{1, 1, 0}, {0, 0, 0}, {0, 0, 0}}, false},
		{[][]int{{0}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, removeOnes(tc.grid))
		})
	}
}

func removeOnes(grid [][]int) bool {
	// Flip each row where the first column contains a 1
	for i := range grid {
		if grid[i][0] == 1 {
			for j := range grid[i] {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	// Flip each column where the first row contains a 1
	for j := range grid[0] {
		if grid[0][j] == 1 {
			for i := range grid {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	// If there are any 1s, the grid cannot be flipped
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

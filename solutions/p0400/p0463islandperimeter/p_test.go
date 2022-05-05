package p0463islandperimeter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_islandPerimeter(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, 16},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, islandPerimeter(tc.grid))
		})
	}
}

func islandPerimeter(grid [][]int) int {
	var perim int
	m, n := len(grid), len(grid[0])
	// Convert grid to matrix of bools
	gridBools := make([][]bool, m)
	for i := range gridBools {
		gridBools[i] = make([]bool, n)
		for j := range gridBools[i] {
			gridBools[i][j] = grid[i][j] == 1
		}
	}

	for i := range gridBools {
		for j := range gridBools[i] {
			if grid[i][j] == 0 {
				continue
			}
			// Check perimeter
			if i == 0 || !gridBools[i-1][j] { // Above
				perim++
			}
			if j == 0 || !gridBools[i][j-1] { // Left
				perim++
			}
			if i == m-1 || !gridBools[i+1][j] { // Below
				perim++
			}
			if j == n-1 || !gridBools[i][j+1] { // Right
				perim++
			}
		}
	}
	return perim
}

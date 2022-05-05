package p2088countfertilepyramindsinaland

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPyramids(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}, 2},
		{[][]int{{1, 1, 1}, {1, 1, 1}}, 2},
		{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, 0},
		{[][]int{{1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, 13},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, countPyramids(tc.grid))
		})
	}
}

func countPyramids(grid [][]int) int {
	// To simplify the exercise, count only pyramids pointing to the top, then
	// invert the y-axis and re-count.
	res := countPyramidsUp(grid)
	m := len(grid)
	// For each col
	for col := range grid[0] {
		// Flip cells in y-axis
		for lo, hi := 0, m-1; lo < hi; lo, hi = lo+1, hi-1 {
			grid[lo][col], grid[hi][col] = grid[hi][col], grid[lo][col]
		}
	}
	// Count again
	res += countPyramidsUp(grid)
	return res
}

func countPyramidsUp(grid [][]int) int {
	// For each possible top-position
	// Grow the pyramid in size until either:
	// 1. The width grows out of bounds
	// 2. The base contains a zero.
	//
	// When a pyramid grows in height, then the pyramid of height-1 is also
	// confirmed for the pyramid starting in the position below. This reduces the
	// total number of lookups.
	m, n := len(grid), len(grid[0])
	valid := make([][][]bool, m)
	for i := range grid {
		valid[i] = make([][]bool, n)
		for j := range grid[i] {
			valid[i][j] = make([]bool, m)
		}
	}
	var count int
	for col := 1; col < n-1; col++ {
		for row := 0; row < m-1; row++ {
			if grid[row][col] == 0 {
				continue
			}
			for height := 1; row+height < m; height++ {
				if valid[row][col][height] {
					// Already counted and verified
					continue
				}
				// Attempt to add the base
				l, r := col-height, col+height
				if l < 0 || r >= n {
					break
				}
				for i := l; i <= r; i++ {
					if grid[row+height][i] != 1 {
						goto breakLoop
					}
				}
				// Pyramid is valid, count and add subpyramids as well.
				for i := 0; i < height; i++ {
					valid[row+i][col][height-i] = true
					count++
				}
			}
		breakLoop:
		}
	}
	return count
}

package p2132stampingthegrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_possibleToStamp(t *testing.T) {
	for _, tc := range []struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
		want        bool
	}{
		{[][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 4, 3, true},
		{[][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 2, 2, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, possibleToStamp(tc.grid, tc.stampHeight, tc.stampWidth))
		})
	}
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	// Create a 2d prefix sum from the grid.
	// For 1D the sum of a range [i,j) is presum[i]-presum[j],
	// For 2D the range between upper left corner [i1, j1) and bottom right
	// [i2, j2) is presum[i2,j2] - presum[i2][j1] - presum[i1][j2] + presum[i1][j1]

	calcPresum := func(a [][]int) [][]int {
		n := len(a)
		m := len(a[0])
		presum := make([][]int, n+1)
		for i := range presum {
			presum[i] = make([]int, m+1)
		}
		for i := range a {
			for j := range a[i] {
				presum[i+1][j+1] = presum[i][j+1] + presum[i+1][j] - presum[i][j] + a[i][j]
			}
		}
		return presum
	}

	pre := calcPresum(grid)
	n := len(grid)
	m := len(grid[0])
	delta := make([][]int, n+1)
	for i := range delta {
		delta[i] = make([]int, m+1)
	}
	for i := 0; i <= n-stampHeight; i++ {
		for j := 0; j <= m-stampWidth; j++ {
			sum := pre[i+stampHeight][j+stampWidth] - pre[i][j+stampWidth] -
				pre[i+stampHeight][j] + pre[i][j]
			if sum == 0 {
				delta[i][j]++
				delta[i+stampHeight][j]--
				delta[i][j+stampWidth]--
				delta[i+stampHeight][j+stampWidth]++
			}
		}
	}
	pre2 := calcPresum(delta)
	for i := range grid {
		for j := range grid[i] {
			if pre2[i+1][j+1] == 0 && grid[i][j] != 1 {
				return false
			}
		}
	}
	return true
}

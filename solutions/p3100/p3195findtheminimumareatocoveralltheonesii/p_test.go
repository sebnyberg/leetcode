package p3195findtheminimumareatocoveralltheonesii

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumArea(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[0,1,0],[1,0,1]]"), 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minimumArea(tc.grid))
		})
	}
}

func minimumArea(grid [][]int) int {
	minX := math.MaxInt32
	minY := math.MaxInt32
	var maxX, maxY int
	for i := range grid {
		for j, val := range grid[i] {
			if val == 1 {
				minX = min(minX, j)
				maxX = max(maxX, j)
				minY = min(minY, i)
				maxY = max(maxY, i)
			}
		}
	}
	return (maxX - minX + 1) * (maxY - minY + 1)
}

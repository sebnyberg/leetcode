package p1591strangeprinterii

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPrintable(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want bool
	}{
		{[][]int{{1, 2, 1}, {2, 1, 2}, {1, 2, 1}}, false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isPrintable(tc.grid))
		})
	}
}

func isPrintable(grid [][]int) bool {
	// Just find rectangles in reverse. The last stamp of a certain color must
	// be a rectangle. When a rectangle is found, mark its are with wildcards.
	//
	// This is ugly but it aint too bad.
	//
	type boundingBox struct {
		upper int
		left  int
		right int
		lower int
	}
	colors := make(map[int]*boundingBox)
	for i := range grid {
		for j, v := range grid[i] {
			if _, exists := colors[v]; !exists {
				colors[v] = &boundingBox{
					upper: math.MinInt32,
					lower: math.MaxInt32,
					left:  math.MaxInt32,
					right: math.MinInt32,
				}
			}
			colors[v].left = min(colors[v].left, j)
			colors[v].right = max(colors[v].right, j)
			colors[v].upper = max(colors[v].upper, i)
			colors[v].lower = min(colors[v].lower, i)
		}
	}
	ok := func(color int, bb *boundingBox) bool {
		for i := bb.lower; i <= bb.upper; i++ {
			for j := bb.left; j <= bb.right; j++ {
				if grid[i][j] > 0 && grid[i][j] != color {
					return false
				}
			}
		}
		return true
	}
colorLoop:
	for len(colors) > 0 {
		for color, bb := range colors {
			if !ok(color, bb) {
				continue
			}
			for i := bb.lower; i <= bb.upper; i++ {
				for j := bb.left; j <= bb.right; j++ {
					grid[i][j] = 0
				}
			}
			delete(colors, color)
			continue colorLoop
		}

		return false
	}
	return true
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

package p0850rectanglearea2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rectangleArea(t *testing.T) {
	for _, tc := range []struct {
		rectangles [][]int
		want       int
	}{
		{[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}, 6},
		{[][]int{{0, 0, 1000000000, 1000000000}}, 49},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rectangles), func(t *testing.T) {
			require.Equal(t, tc.want, rectangleArea(tc.rectangles))
		})
	}
}

const mod = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	// Path compress x- and y-axis
	xUnique := make(map[int]struct{})
	yUnique := make(map[int]struct{})
	for _, rect := range rectangles {
		xUnique[rect[0]] = struct{}{}
		xUnique[rect[2]] = struct{}{}
		yUnique[rect[1]] = struct{}{}
		yUnique[rect[3]] = struct{}{}
	}

	// Create list of x / y values (for sorting)
	xVals := make([]int, 0, len(xUnique))
	for x := range xUnique {
		xVals = append(xVals, x)
	}
	yVals := make([]int, 0, len(yUnique))
	for y := range yUnique {
		yVals = append(yVals, y)
	}

	// Sort
	sort.Ints(xVals)
	sort.Ints(yVals)

	// Create translation map
	xTrans := make(map[int]int, len(xVals))
	yTrans := make(map[int]int, len(yVals))
	for i, x := range xVals {
		xTrans[x] = i
	}
	for i, y := range yVals {
		yTrans[y] = i
	}

	// Count number of rectangles for each position in the compressed grid
	grid := make([][]int, len(yVals))
	for y := range grid {
		grid[y] = make([]int, len(xVals))
	}
	for _, rect := range rectangles {
		xStart, xEnd := xTrans[rect[0]], xTrans[rect[2]]
		yStart, yEnd := yTrans[rect[1]], yTrans[rect[3]]
		for y := yStart; y < yEnd; y++ {
			grid[y][xStart]++
			grid[y][xEnd]--
		}
	}

	// Calculate area
	var area int
	for yIdx := 0; yIdx < len(yVals)-1; yIdx++ {
		// Calculate width of rectangles covering current line sweep
		yDelta := yVals[yIdx+1] - yVals[yIdx]
		var nrect int
		var start int
		for xIdx := 0; xIdx < len(xVals); xIdx++ {
			if grid[yIdx][xIdx] == 0 {
				continue
			}
			if nrect == 0 {
				start = xVals[xIdx] // store start x-position
			}
			nrect += grid[yIdx][xIdx]
			if nrect == 0 {
				area = (area + (xVals[xIdx]-start)*yDelta) % mod
			}
		}
	}

	return area
}

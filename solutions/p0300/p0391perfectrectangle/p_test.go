package p0391perfectrectangle

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isRectangleCover(t *testing.T) {
	for _, tc := range []struct {
		rectangles [][]int
		want       bool
	}{
		{[][]int{{0, 0, 1, 1}, {0, 0, 2, 1}, {1, 0, 2, 1}, {0, 2, 2, 3}}, false},
		{[][]int{{0, 0, 1, 1}, {0, 2, 1, 3}, {1, 1, 2, 2}, {2, 0, 3, 1}, {2, 2, 3, 3}, {1, 0, 2, 3}, {0, 1, 3, 2}}, false},
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}, true},
		{[][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}, false},
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}}, false},
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rectangles), func(t *testing.T) {
			require.Equal(t, tc.want, isRectangleCover(tc.rectangles))
		})
	}
}

func isRectangleCover(rectangles [][]int) bool {
	// Based on discussion post on the forum.
	// Assuming no overlap or gap, any corner with an odd count must be a global
	// corner. The total size of all rectangles should sum up to the size of the
	// rectangle formed by those odd count corners.
	type corner struct{ x, y int }
	corners := make(map[corner]struct{}, 4)
	var expectedSize int
	for _, rect := range rectangles {
		x, y, X, Y := rect[0], rect[1], rect[2], rect[3]
		newCorners := []corner{{x, Y}, {X, Y}, {x, y}, {X, y}}
		expectedSize += (Y - y) * (X - x)
		for _, newCorner := range newCorners {
			if _, exists := corners[newCorner]; exists {
				delete(corners, newCorner)
			} else {
				corners[newCorner] = struct{}{}
			}
		}
	}
	if len(corners) != 4 {
		return false
	}
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	for corner := range corners {
		minX, minY = min(minX, corner.x), min(minY, corner.y)
		maxX, maxY = max(maxX, corner.x), max(maxY, corner.y)
	}
	return (maxY-minY)*(maxX-minX) == expectedSize
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

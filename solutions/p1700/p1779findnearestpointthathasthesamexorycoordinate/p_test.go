package p1779findnearestpointthathasthesamexorycoordinate

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nearestValidPoint(t *testing.T) {
	for _, tc := range []struct {
		x      int
		y      int
		points [][]int
		want   int
	}{
		{3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, nearestValidPoint(tc.x, tc.y, tc.points))
		})
	}
}

func nearestValidPoint(x int, y int, points [][]int) int {
	minDist := math.MaxInt32
	minIdx := -1
	for i, point := range points {
		if point[0] == x {
			if d := abs(point[1] - y); d < minDist {
				minDist = d
				minIdx = i
			}
			continue
		}
		if point[1] == y {
			if d := abs(point[0] - x); d < minDist {
				minDist = d
				minIdx = i
			}
		}
	}
	return minIdx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

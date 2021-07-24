package p0356linereflection

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isReflected(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   bool
	}{
		{[][]int{{1, 1}, {-1, 1}}, true},
		{[][]int{{1, 1}, {-1, -1}}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, isReflected(tc.points))
		})
	}
}

func isReflected(points [][]int) bool {
	maxX := math.MinInt32
	minX := math.MaxInt32
	pointMap := make(map[[2]int]struct{})
	for _, p := range points {
		maxX = max(maxX, p[0])
		minX = min(minX, p[0])
		pointMap[[2]int{p[0], p[1]}] = struct{}{}
	}
	center := (float64(minX) + float64(maxX)) / 2
	for _, p := range points {
		reflectedX := int(math.Round(center + (center - float64(p[0]))))
		if _, exists := pointMap[[2]int{reflectedX, p[1]}]; !exists {
			return false
		}
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

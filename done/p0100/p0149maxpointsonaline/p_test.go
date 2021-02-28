package p0149maxpointsonaline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPoints(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{[][]int{{0, 0}}, 1},
		{[][]int{{-4, -4}, {-8, -582}, {-3, 3}, {-9, -651}, {9, 591}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, maxPoints(tc.points))
		})
	}
}

func maxPoints(points [][]int) int {
	maxPointsOnALine := 1
	n := len(points)

	// For each point in the list
	for i, first := range points {

		// Consider each other point and count the number of points sharing slopes
		// Avoid rounding error mismatches by storing the y/x integers in their
		// smallest form (divide by GCD)
		// e.g. 18/4 -> 9/2
		slopeCount := make(map[[2]int]int, n-i)
		for _, second := range points[i+1:] {

			// Edge case: first and second points share x value, infinite slope
			if first[0] == second[0] {
				slopeCount[[2]int{0, 0}]++
				maxPointsOnALine = max(maxPointsOnALine, slopeCount[[2]int{0, 0}]+1)
				continue
			}

			// Divide y/x with GCD and add it to the slope count
			xDiff := first[0] - second[0]
			yDiff := first[1] - second[1]
			dividend := gcd(yDiff, xDiff)
			xDiff /= dividend
			yDiff /= dividend
			slopeCount[[2]int{xDiff, yDiff}]++
			maxPointsOnALine = max(maxPointsOnALine, slopeCount[[2]int{xDiff, yDiff}]+1)
		}
	}

	return maxPointsOnALine
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(m int, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

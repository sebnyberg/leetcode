package p1937maximumnumberofpointswithcost

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPoints(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int64
	}{
		{[][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}, 9},
		{[][]int{{1, 5}, {2, 3}, {4, 2}}, 11},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, maxPoints(tc.points))
		})
	}
}

func maxPoints(points [][]int) int64 {
	// For each row in points,
	// Calculate the maximum possible value for each position in that row given
	// the provided constraints.
	n := len(points[0])
	// prev[i] denotes the max points for position i in the previous row
	val := make([]int, n)

	for i := range points {
		for j := range points[i] {
			val[j] += points[i][j]
		}

		// Two-pass max value collector
		for j := 1; j < n; j++ {
			val[j] = max(val[j], val[j-1]-1)
		}
		for j := n - 2; j >= 0; j-- {
			val[j] = max(val[j], val[j+1]-1)
		}
	}
	var maxPoints int
	for _, v := range val {
		maxPoints = max(maxPoints, v)
	}
	return int64(maxPoints)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

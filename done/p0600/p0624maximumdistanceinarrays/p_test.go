package p0624maximumdistanceinarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDistance(t *testing.T) {
	for _, tc := range []struct {
		arrays [][]int
		want   int
	}{
		{[][]int{{-8, -7, -7, -5, 1, 1, 3, 4}, {-2}, {-10, -10, -7, 0, 1, 3}, {2}}, 14},
		{[][]int{{-8}, {-3, 1, 4}, {-2, -1, 0, 2}}, 12},
		{[][]int{{-1, 1}, {-3, 1, 4}, {-2, -1, 0, 2}}, 6},
		{[][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4},
		{[][]int{{1}, {1}}, 0},
		{[][]int{{1}, {2}}, 1},
		{[][]int{{1, 4}, {0, 5}}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arrays), func(t *testing.T) {
			require.Equal(t, tc.want, maxDistance(tc.arrays))
		})
	}
}

func maxDistance(arrays [][]int) int {
	minVal, maxVal := arrays[0][0], arrays[0][len(arrays[0])-1]
	var maxDist int
	for i := 1; i < len(arrays); i++ {
		maxDist = max(maxDist, max(maxVal-arrays[i][0], arrays[i][len(arrays[i])-1]-minVal))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][len(arrays[i])-1])
	}
	return maxDist
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

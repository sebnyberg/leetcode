package p0256painthouse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for _, tc := range []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}, 10},
		{[][]int{{7, 6, 2}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.costs), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.costs))
		})
	}
}

func minCost(costs [][]int) int {
	colorCost := make([]int, 3)
	for _, cost := range costs {
		colorCost[0], colorCost[1], colorCost[2] =
			min(colorCost[1], colorCost[2])+cost[0],
			min(colorCost[0], colorCost[2])+cost[1],
			min(colorCost[0], colorCost[1])+cost[2]
	}
	return min(colorCost[0], min(colorCost[1], colorCost[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

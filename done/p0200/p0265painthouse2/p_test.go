package p0265painthouse2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostII(t *testing.T) {
	for _, tc := range []struct {
		costs [][]int
		want  int
	}{
		{[][]int{
			{20, 19, 11, 13, 12, 16, 16, 17, 15, 9, 5, 18},
			{3, 8, 15, 17, 19, 8, 18, 3, 11, 6, 7, 12},
			{15, 4, 11, 1, 18, 2, 10, 9, 3, 6, 4, 15},
		}, 9},
		{[][]int{{1, 5, 3}, {2, 9, 4}}, 5},
		{[][]int{{1, 3}, {2, 4}}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.costs), func(t *testing.T) {
			require.Equal(t, tc.want, minCostII(tc.costs))
		})
	}
}

func minCostII(costs [][]int) int {
	k := len(costs[0])
	dp := make([]int, k)

	var (
		firstMin    = 0
		firstMinIdx = -1
		secondMin   = 0
	)
	for _, cost := range costs {
		for j, colorCost := range cost {
			if j == firstMinIdx {
				dp[j] = secondMin + colorCost
			} else {
				dp[j] = firstMin + colorCost
			}
		}
		firstMin, firstMinIdx, secondMin = math.MaxInt32, -1, math.MaxInt32
		for j, colorCost := range dp {
			if colorCost < firstMin {
				secondMin = firstMin
				firstMin, firstMinIdx = colorCost, j
			} else if colorCost < secondMin {
				secondMin = colorCost
			}
		}
	}
	return firstMin
}

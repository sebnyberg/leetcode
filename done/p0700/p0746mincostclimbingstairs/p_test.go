package p0746mincostclimbingstairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostClimbingStairs(t *testing.T) {
	for _, tc := range []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cost), func(t *testing.T) {
			require.Equal(t, tc.want, minCostClimbingStairs(tc.cost))
		})
	}
}

func minCostClimbingStairs(cost []int) int {
	// dp keeps track of the min cost from 0 steps ago and 1 steps ago
	dp := [2]int{0, cost[0]}
	for i := range cost[1:] {
		dp[0], dp[1] = dp[1], min(dp[0], dp[1])+cost[i+1]
	}

	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

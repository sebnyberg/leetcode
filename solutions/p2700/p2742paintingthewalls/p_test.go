package p2742paintingthewalls

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_paintWalls(t *testing.T) {
	for i, tc := range []struct {
		cost []int
		time []int
		want int
	}{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 2}, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, paintWalls(tc.cost, tc.time))
		})
	}
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := range cost {
		for j := n - 1; j >= 0; j-- {
			k := min(n, j+1+time[i])
			dp[k] = min(dp[k], dp[j]+cost[i])
		}
	}
	return dp[n]
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

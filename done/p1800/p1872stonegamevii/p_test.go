package p1872stonegamevii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stoneGameVIII(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		want   int
	}{
		{[]int{-1, 2, -3, 4, -5}, 5},
		{[]int{7, -6, 5, 10, 5, -2, -6}, 13},
		{[]int{-10, -12}, -22},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, stoneGameVIII(tc.stones))
		})
	}
}

func stoneGameVIII(stones []int) int {
	// Alice wants to maximize the score, thus she seeks the greaters possible
	// prefix sum in the stones array
	n := len(stones)
	sums := make([]int, n)
	sums[0] = stones[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + stones[i]
	}
	dp := make([]int, n)
	dp[n-1] = sums[n-1]
	for i := n - 2; i >= 1; i-- {
		dp[i] = max(dp[i+1], sums[i]-dp[i+1])
	}
	return dp[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

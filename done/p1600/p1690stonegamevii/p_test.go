package p1690stonegamevii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stoneGameVII(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		want   int
	}{
		{[]int{5, 3, 1, 4, 2}, 6},
		{[]int{7, 90, 5, 1, 100, 10, 10, 2}, 122},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, stoneGameVII(tc.stones))
		})
	}
}

func stoneGameVII(stones []int) int {
	n := len(stones)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + stones[i]
	}

	for k := 2; k <= n; k++ {
		for start := 0; start+k-1 < n; start++ {
			end := start + k - 1
			// Remove either the left stone, summing the remainder,
			// or the right stone, summing the remainder
			scoreRemoveFirst := presum[end+1] - presum[start+1]
			scoreRemoveLast := presum[end] - presum[start]

			dp[start][end] = max(
				scoreRemoveFirst-dp[start+1][end], // remove left
				scoreRemoveLast-dp[start][end-1],  // remove right
			)
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0410splitarraylargestsum

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		m    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 15},
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, splitArray(tc.nums, tc.m))
		})
	}
}

func splitArray(nums []int, m int) int {
	// dp[i][k] = max sum when splitting [0..i] into k parts
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			// Guard against one-off issues
			dp[i][j] = math.MaxInt32
		}
	}
	presum := make([]int, n+1)
	for i, num := range nums {
		presum[i+1] = presum[i] + num
		dp[i][1] = presum[i]
	}
	dp[n][1] = presum[n]
	for end := 1; end <= n; end++ {
		for k := 2; k <= m && k <= end; k++ {
			// The minimum value between [0..end] is found by trying all possible
			// splits which divide the left side in k-1 sums, and right in one sum.
			for mid := 1; mid < end; mid++ {
				dp[end][k] = min(
					dp[end][k],
					max(dp[mid][k-1], presum[end]-presum[mid]),
				)
			}
		}
	}
	return dp[n][m]
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

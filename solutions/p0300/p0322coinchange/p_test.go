package p0322coinchange

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_coinChange(t *testing.T) {
	for _, tc := range []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.coins, tc.amount), func(t *testing.T) {
			require.Equal(t, tc.want, coinChange(tc.coins, tc.amount))
		})
	}
}

func coinChange(coins []int, amount int) int {
	// Since amount is small, we may DP to find the result.
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0589coinchange

import (
	"fmt"
	"math"
	"sort"
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
	sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				break
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] < math.MaxInt32 {
		return dp[amount]
	}
	return -1
}

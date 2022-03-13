package p0518coinchange2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_change(t *testing.T) {
	for _, tc := range []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.amount), func(t *testing.T) {
			require.Equal(t, tc.want, change(tc.amount, tc.coins))
		})
	}
}

func change(amount int, coins []int) int {
	var dp [5001]int
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

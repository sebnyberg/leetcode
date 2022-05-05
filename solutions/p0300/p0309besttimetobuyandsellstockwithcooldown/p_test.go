package p0309besttimetobuyandsellstockwithcooldown

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		want   int
	}{
		{[]int{2, 1}, 0},
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.prices))
		})
	}
}

func maxProfit(prices []int) int {
	sell, buy, rest := math.MinInt32, math.MinInt32, 0

	// At a given time t
	// If you wish to buy, the total value will be the max sell value two steps ago
	// If you wish to sell, the total value will become the previous buy value + price
	for i := 0; i < len(prices); i++ {
		sell, buy, rest = buy+prices[i],
			max(buy, rest-prices[i]),
			max(rest, sell)
	}
	return max(sell, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

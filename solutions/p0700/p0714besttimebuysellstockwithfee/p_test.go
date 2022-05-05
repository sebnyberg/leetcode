package p0714besttimebuysellstockwithfee

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		fee    int
		want   int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.prices, tc.fee), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.prices, tc.fee))
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	buy, sell := -prices[0]-fee, 0
	for _, price := range prices[1:] {
		buy, sell = max(buy, sell-price-fee), max(sell, buy+price)
	}
	return max(buy, sell)
}

package p0121buyandsellstock

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
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.prices))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maximumProfit := math.MinInt32
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maximumProfit = max(maximumProfit, price-minPrice)
	}
	return maximumProfit
}

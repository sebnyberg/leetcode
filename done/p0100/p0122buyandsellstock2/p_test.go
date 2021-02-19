package p0122buyandsellstock2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.prices))
		})
	}
}

func maxProfit(prices []int) (profit int) {
	if len(prices) <= 1 {
		return 0
	}

	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

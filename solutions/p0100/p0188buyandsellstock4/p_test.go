package p0188buyandsellstock4

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	for _, tc := range []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.k, tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.k, tc.prices))
		})
	}
}

func maxProfit(k int, prices []int) int {
	// For a given day, you can either hold, sell, or buy.
	//
	// Since prices are non-negative, the optimal result will be given by selling
	// 0 to k times.
	//
	// Since the optimal is given by selling, it doesn't matter whether we
	// consider selling or buying to be a "transaction increaser". Here I consider
	// buying to be a transaction increaser, but it doesn't really matter.
	//
	// For a given price, we must consider all possible states: holding with k
	// transactions, and not holding with k transactions. We can then choose
	// either to sell when we hold, buy so that we start to hold (and increase the
	// number of transactions), or not do anything.
	//
	// Initially I wrote this as a curr + next state, but by going from end=k to
	// start I could use only a single state array.
	var curr [101][2]int
	const (
		notHolding = 0
		holding    = 1
	)
	for i := range curr {
		curr[i][notHolding] = math.MinInt32
		curr[i][holding] = math.MinInt32
	}
	curr[0][notHolding] = 0
	var maxProfit int

	for _, v := range prices {
		for j := k; j >= 1; j-- {
			curr[j][notHolding] = max(curr[j][notHolding], curr[j][holding]+v)
			curr[j][holding] = max(curr[j][holding], curr[j-1][notHolding]-v)
			maxProfit = max(maxProfit, curr[j][notHolding])
		}
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

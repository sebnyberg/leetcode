package p0879profitableschemes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_profitableSchemes(t *testing.T) {
	for i, tc := range []struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		want      int
	}{
		{64, 0, []int{80, 40}, []int{88, 88}, 2},
		{10, 5, []int{2, 3, 5}, []int{6, 7, 8}, 7},
		{5, 3, []int{2, 2}, []int{2, 3}, 2},
		{10, 5, []int{2, 3, 5}, []int{6, 7, 8}, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, profitableSchemes(tc.n, tc.minProfit, tc.group, tc.profit))
		})
	}
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	// This is a typical dp problem.
	// Each crime can either contribute or not contribute to all prior possible
	// series of crime, given that there are enough people.
	//
	// This tells us that we need to keep track of all prior possible levels of
	// profit, and how many people remain for those levels of profit.
	//
	// This is O(minProfit*n*group.length)
	//
	var dp [101][101]int
	dp[n][0] = 1
	const mod = 1e9 + 7
	for i := range group {
		for m := 0; m <= n-group[i]; m++ {
			prevGroup := m + group[i]
			for prevProfit := minProfit; prevProfit >= 0; prevProfit-- {
				p := min(minProfit, prevProfit+profit[i])
				dp[m][p] += dp[prevGroup][prevProfit]
				dp[m][p] %= mod
			}
		}
	}
	var res int
	for m := 0; m <= n; m++ {
		res = (res + dp[m][minProfit]) % mod
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

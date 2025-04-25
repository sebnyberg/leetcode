package p0375guessnumberhigherorlower2

import (
	"fmt"
	"math"
	"testing"
)

func Test_getMoneyAmount(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{200, 952},
		{10, 16},
		{1, 0},
		{2, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, getMoneyAmount(tc.n))
		})
	}
}

func getMoneyAmount(n int) int {
	// Given the interval [1, n], the number of guesses in a regular binary search
	// would be log(n). However, there is now a cost function involved, where we
	// wish to minimize the worst-case cost.

	// Use dynamic programming to determine the minimum worst-case cost per
	// interval length. dp[i][j] = min worst-case cost from i to j (exclusive)
	dp := make([][]uint16, n+2)
	for i := range dp {
		dp[i] = make([]uint16, n+2)
	}
	// The cost for one option is zero
	// Iterate over each interval length
	for k := 2; k <= n; k++ {
		for i := 1; i <= n-k+1; i++ {
			// The best guess minimizes the max cost of intervals to the left and
			// right of the chosen position
			minResult := uint16(math.MaxUint16)
			for j := 0; j < k; j++ {
				left := dp[i][i+j]
				right := dp[i+j+1][i+k]
				minResult = min(minResult, uint16(i+j)+max(left, right))
			}
			dp[i][i+k] = minResult
		}
	}
	return int(dp[1][n+1])
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

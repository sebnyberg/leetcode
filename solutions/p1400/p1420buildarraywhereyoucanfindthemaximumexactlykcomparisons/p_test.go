package p1420buildarraywhereyoucanfindthemaximumexactlykcomparisons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numOfArrays(t *testing.T) {
	for i, tc := range []struct {
		n    int
		m    int
		k    int
		want int
	}{
		{2, 4, 2, 6},
		{2, 3, 1, 6},
		{5, 2, 3, 0},
		{9, 1, 1, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numOfArrays(tc.n, tc.m, tc.k))
		})
	}
}

const mod = 1e9 + 7

const N = 50
const M = 100
const K = 50

func numOfArrays(n int, m int, k int) int {
	// It very much looks like a DP-problem
	// For each index we may either increase the current maximum, or not.
	// If we do, we increase search_cost by 1
	// So we need to know:
	//
	// 1. Current search score
	// 2. Current maximum value
	// 3. Current index
	//
	// dp[i][k][m] = at number i, number of ways to have a total search score of
	// k given that the current maximum value is m
	var dp [N][K + 1][M + 1]int

	// Initial round
	for x := 1; x <= m; x++ {
		dp[0][1][x] = 1
	}

	for i := 1; i < n; i++ {
		// For each prior maximum value
		for x := 1; x <= m; x++ {
			// And new value
			for y := 1; y <= m; y++ {
				mm := max(x, y)
				// If the new value is greater than the prior max, then the
				// search score is increased
				var s int
				if y > x {
					s++
				}
				// For each possible new search score
				for kk := 1; kk+s <= k; kk++ {
					// Add number of ways to reach the prior search score
					dp[i][kk+s][mm] = (dp[i][kk+s][mm] + dp[i-1][kk][x]) % mod
				}
			}
		}
	}
	var res int
	for _, x := range dp[n-1][k] {
		res = (res + x) % mod
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

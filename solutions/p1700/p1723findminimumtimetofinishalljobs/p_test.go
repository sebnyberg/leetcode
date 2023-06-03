package p1723findminimumtimetofinishalljobs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTimeRequired(t *testing.T) {
	for i, tc := range []struct {
		jobs []int
		k    int
		want int
	}{
		{[]int{3, 2, 3}, 3, 3},
		{[]int{1, 2, 4, 7, 8}, 2, 11},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTimeRequired(tc.jobs, tc.k))
		})
	}
}

func minimumTimeRequired(jobs []int, k int) int {
	// The basic idea here is to cache all possible sums, then try all possible
	// combinations of selections. It's brute-force + memoization.
	n := len(jobs)
	m := (1 << n) - 1
	dp := make([][]int, k+1)
	for i := 2; i <= k; i++ {
		dp[i] = make([]int, m+1)
	}

	sums := make([]int, m+1)
	for x := 1; x <= m; x++ {
		for i := 0; i < n; i++ {
			if x&(1<<i) > 0 {
				sums[x] += jobs[i]
			}
		}
	}

	dp[1] = sums
	for i := 2; i <= k; i++ {
		for x := 1; x <= m; x++ {
			dp[i][x] = dp[i-1][x] // base case
			for s := x; s > 0; s = (s - 1) & x {
				dp[i][x] = min(dp[i][x], max(sums[s], dp[i-1][x^s]))
			}
		}
	}
	return dp[k][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

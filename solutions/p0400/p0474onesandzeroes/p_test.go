package p0474onesandzeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxForm(t *testing.T) {
	for _, tc := range []struct {
		strs []string
		m    int
		n    int
		want int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxForm(tc.strs, tc.m, tc.n))
		})
	}
}

// m = max number of zeroes
// n = max number of ones
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// For each string
	for _, str := range strs {
		// Count ones and zeroes
		ones, zeroes := countOnesAndZeroes(str)
		for i := m; i >= zeroes; i-- {
			for j := n; j >= ones; j-- {
				// For each location in the mxn DP array,
				// The max strings from the input for a given position is either
				// 1. the max at a distance reachable with the currrent string + 1
				// 2. the previous value
				dp[i][j] = max(dp[i-zeroes][j-ones]+1, dp[i][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countOnesAndZeroes(s string) (int, int) {
	ones, zeroes := 0, 0
	for _, r := range s {
		if r == '1' {
			ones++
		} else {
			zeroes++
		}
	}
	return ones, zeroes
}

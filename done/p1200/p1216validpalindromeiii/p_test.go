package p1216validpalindromeiii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValidPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want bool
	}{
		{"abcdeca", 2, true},
		{"abbababa", 1, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isValidPalindrome(tc.s, tc.k))
		})
	}
}

func isValidPalindrome(s string, k int) bool {
	// Match s and a reversed version of s
	rev := revStr(s)

	// Construct a DP output. Empty strings are matches.
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = j
			}
		}
		dp[i][0] = i
	}

	// For each row and col in DP
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// If s[i-1] and rev[j-1] is equal, then we can move diagonally in the
			// comparison.
			if s[j-1] == rev[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// Else, this matching is invalid, and so the best solution is to skip
				// this letter + have a prior valid matching, either left or above in
				// the dp matrix.
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}

			if dp[i][j] <= k && i+j-k >= n-k {
				return true
			}
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func revStr(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

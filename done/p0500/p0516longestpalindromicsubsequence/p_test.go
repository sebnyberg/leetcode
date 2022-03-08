package p0516longestpalindromicsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindromeSubseq(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"a", 1},
		{"bbbab", 4},
		{"cbbd", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestPalindromeSubseq(tc.s))
		})
	}
}

func longestPalindromeSubseq(s string) int {
	var dp [1001][1001]int
	n := len(s)
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 1; r < n; r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1] + 2
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

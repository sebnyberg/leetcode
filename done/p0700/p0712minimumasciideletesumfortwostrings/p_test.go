package p0712minimumasciideletesumfortwostrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDeleteSum(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDeleteSum(tc.s1, tc.s2))
		})
	}
}

func minimumDeleteSum(s1 string, s2 string) int {
	// Looks like a DP exercise
	// We may either match two characters at zero cost
	// Or delete one from one at ASCII cost for that character,
	// or vice versa.
	n1, n2 := len(s1), len(s2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= n2; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = min(
				dp[i-1][j]+int(s1[i-1]), // Delete from first string
				dp[i][j-1]+int(s2[j-1]), // Delete from second string
			)

			// If there's a match, we may use it
			if s1[i-1] == s2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			}
		}
	}

	return dp[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

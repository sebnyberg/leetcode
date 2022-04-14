package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strangePrinter(t *testing.T) {
	type testCase struct {
		s    string
		want int
	}

	testCases := []testCase{
		{"tbgtgb", 4},
		{"aaabbb", 2},
		{"aba", 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase %v", i), func(t *testing.T) {
			require.Equal(t, tc.want, strangePrinter(tc.s))
		})
	}
}

func strangePrinter(s string) int {
	n := len(s)
	var dp [101][101]uint8
	for i := range dp {
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for d := 1; i+d < n; d++ {
			j := i + d
			if d == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 2
				}
				continue
			}

			dp[i][j] = 101
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
			if s[i] == s[j] {
				dp[i][j]--
			}
		}
	}

	return int(dp[0][n-1])
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

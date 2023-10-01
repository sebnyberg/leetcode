package p2707extracharactersinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minExtraChar(t *testing.T) {
	for i, tc := range []struct {
		s          string
		dictionary []string
		want       int
	}{
		{"leetscode", []string{"leet", "code", "leetcode"}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minExtraChar(tc.s, tc.dictionary))
		})
	}
}

func minExtraChar(s string, dictionary []string) int {
	// Consider a single position - does it matter how the previous substrings
	// were formed, given that they all had the same amount of skipped
	// characters? No.
	//
	// This means that we could consider each position independently. The same
	// logic also goes for forward-mapping. This tells us that we can use DP to
	// find the solution.
	//
	// For each position in the string, from end to start, the current best
	// solution is given by combining all words in the dictionary with the
	// current position (if applicable) and adding that to the best result from
	// where the word ends.
	//
	n := len(s)
	dp := make([]int, n+1)
	for i := range s {
		dp[i] = n - i
	}
	for i := n - 1; i >= 0; i-- {
		for _, w := range dictionary {
			if i+len(w) > n {
				continue
			}
			if s[i:i+len(w)] != w {
				continue
			}
			dp[i] = min(dp[i], dp[i+len(w)])
		}
		dp[i] = min(dp[i], dp[i+1]+1)
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

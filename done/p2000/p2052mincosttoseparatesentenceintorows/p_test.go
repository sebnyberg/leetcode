package p2052mincosttoseparatesentenceintorows

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCost(t *testing.T) {
	for _, tc := range []struct {
		sentence string
		k        int
		want     int
	}{
		{"aa aaa aa", 6, 0},
		{"aa aaa aa", 10, 0},
		{"apples and bananas taste great", 7, 21},
		{"i love leetcode", 12, 36},
		{"a", 5, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentence), func(t *testing.T) {
			require.Equal(t, tc.want, minimumCost(tc.sentence, tc.k))
		})
	}
}

func minimumCost(sentence string, k int) int {
	// For each word in the sentence
	// For each word prior to that word such that their combined length + spaces
	// between is <= k.
	// Update the minimum value for the current position in dp to be
	// min(dp[i], dp[i-nwords]+cost)
	words := strings.Fields(sentence)
	n := len(words)
	widths := make([]int, n)
	for i := range words {
		widths[i] = len(words[i])
		if widths[i] > k {
			return -1
		}
	}
	dp := make([]int, len(words)+1)
	dp[0] = 0
	for i := range widths {
		tot := widths[i]
		dp[i+1] = dp[i] + (k-tot)*(k-tot)
		for j := i - 1; j >= 0 && tot+widths[j]+1 <= k; j-- {
			tot += widths[j] + 1
			dp[i+1] = min(dp[i+1], dp[j]+(k-tot)*(k-tot))
		}
	}
	minCost := min(dp[n], dp[n-1])
	tot := widths[n-1]
	for j := n - 2; j >= 0 && tot+widths[j]+1 <= k; j-- {
		minCost = min(minCost, dp[j])
		tot += widths[j] + 1
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

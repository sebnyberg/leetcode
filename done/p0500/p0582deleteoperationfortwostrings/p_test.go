package p0582deleteoperationfortwostrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	for _, tc := range []struct {
		word1 string
		word2 string
		want  int
	}{
		{"sea", "eat", 2},
		{"leetcode", "etco", 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word1), func(t *testing.T) {
			require.Equal(t, tc.want, minDistance(tc.word1, tc.word2))
		})
	}
}

func minDistance(word1 string, word2 string) int {
	// We can use DP to create a maximum matching of word1 and word2
	var dp [502][502]int
	n := len(word1)
	m := len(word2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var res int
			if word1[j-1] == word2[i-1] {
				res = 1
			}

			dp[i][j] = max(
				dp[i][j-1], // skip one character in word1
				max(
					dp[i-1][j],       // skip one character in word2
					res+dp[i-1][j-1], // move both forward by 1
				),
			)
		}
	}
	res := len(word1) - dp[m][n]
	res += len(word2) - dp[m][n]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

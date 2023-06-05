package p1771maximizepalindromelengthfromsubsequences

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	for i, tc := range []struct {
		word1 string
		word2 string
		want  int
	}{
		{"febeeb", "d", 0},
		{"cacb", "cbba", 5},
		{"ab", "ab", 3},
		{"aa", "bb", 0},
		{"afaaadacb", "ca", 6},
		{"abdbdddb", "bbcdeedbdfccdeabaaaadcdecbdbfbaadbfeaedc", 27},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestPalindrome(tc.word1, tc.word2))
		})
	}
}

func longestPalindrome(word1 string, word2 string) int {
	// This problem is a finicky version of longest match between two strings.
	//
	// The solution is to match the concatenation of the two words with the
	// reverse of the concatenation.
	//
	// However, many of the matchings are not valid. For example, we cannot
	// leave the boundary of word1 or word2 without having matched at least one
	// character.
	//
	// Also, any matching that uses indices beyond the length of word1+word2 are
	// invalid.
	//
	// Also, once a matching has been found, unless the final characters are
	// next to each other, we could add a letter in the middle of the matching
	// to increase the total size of the palindrome.
	//
	// Note that this can be done with 1D-DP, but I can't be arsed to re-write
	// it.
	s := word1 + word2
	buf := []byte(s)
	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	t := string(buf)

	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+n+1)
	for i := range dp {
		dp[i] = make([]int, m+n+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = 0
	}

	var res int

	for i := 1; i <= m+n; i++ {
		for j := 1; j <= m+n; j++ {
			if i+j > m+n {
				break
			}

			if i-1 == len(word1) {
				// If crossing from word1 to word2, ensure that at least one
				// character has matched so far
				if dp[i-1][j] > 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j])
				}
			} else {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}

			if j-1 == len(word2) {
				// If crossing from word2 to word1 in the reverse matching,
				// ensure that at least one character has matched so far
				if dp[i][j-1] > 0 {
					dp[i][j] = max(dp[i][j], dp[i][j-1])
				}
			} else {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}

			var diag int
			if s[i-1] == t[j-1] {
				diag = 1
			}
			if i-1 == len(word1) || j-1 == len(word2) {
				// Same for the diagonal
				if dp[i-1][j-1] > 0 {
					dp[i][j] = max(dp[i][j], diag+dp[i-1][j-1])
				}
			} else {
				dp[i][j] = max(dp[i][j], diag+dp[i-1][j-1])
			}

			var middleChar int
			if i+j < m+n {
				middleChar = 1
			}

			if dp[i][j] > 0 && 2*dp[i][j]+middleChar > res {
				res = 2*dp[i][j] + middleChar
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

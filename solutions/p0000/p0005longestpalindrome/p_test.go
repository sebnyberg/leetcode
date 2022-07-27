package p0005longestpalindromic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	tcs := []struct {
		in   string
		want string
	}{
		{"bb", "bb"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"ccc", "ccc"},
		{"bananas", "anana"},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, longestPalindrome(tc.in))
		})
	}
}

func longestPalindrome(s string) string {
	// There are better approaches here, but since it's only a length 1000 string,
	// I'm going with straight-forward window checking.

	// Time: O(n^2)
	// Space: O(1)

	isPalindrome := func(l, r int) bool {
		for ; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	maxL, maxR, maxLen := 0, 1, 1
	for l := 0; l < len(s)-1; l++ {
		for r := l + 1; r <= len(s); r++ {
			if r-l <= maxLen || !isPalindrome(l, r-1) {
				continue
			}
			maxLen = r - l
			maxL = l
			maxR = r
		}
	}

	return s[maxL:maxR]
}

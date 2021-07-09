package p0393utf8validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSubstring(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"bchhbbdefghiaaacb", 3, 3},
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
		{"abcccde", 2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestSubstring(tc.s, tc.k))
		})
	}
}

func longestSubstring(s string, k int) int {
	// If the longest substring is smaller than len(s), then there must be
	// some characters in s of which there are less then len(s).
	invalidIndices := invalidCharIndices(s, k)
	if len(invalidIndices) == 0 {
		return len(s)
	}

	// Attempt to find substrings that are valid
	var start int
	invalidIndices = append(invalidIndices, len(s))
	var maxLen int
	for _, end := range invalidIndices {
		if end-start >= k {
			maxLen = max(maxLen, longestSubstring(s[start:end], k))
		}
		start = end + 1
	}
	return maxLen
}

func invalidCharIndices(s string, k int) []int {
	var charCount [26]int
	for _, ch := range s {
		charCount[ch-'a']++
	}
	var res []int
	for i, ch := range s {
		if charCount[ch-'a'] < k {
			res = append(res, i)
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

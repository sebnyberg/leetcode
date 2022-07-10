package p0003longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tcs := []struct {
		in      string
		want    int
		wantStr string
	}{
		{"aab", 2, "ab"},
		{"dvdf", 3, "vdf"},
		{"abcabcbb", 3, "abc"},
		{"bbbbb", 1, "b"},
		{"pwwkew", 3, "wke"},
		{"", 0, ""},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLongestSubstring(tc.in))
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	// Greedy approach: move a right pointer until there's a duplicate.
	// Keep track of where each character was last seen so that the left pointer
	// is easy to move. Otherwise we'd have to scan through the string.
	var seenAt [256]int
	for i := range seenAt {
		seenAt[i] = -1
	}
	var res int
	l := -1 // left pointer
	for i, ch := range s {
		l = max(l, seenAt[ch])
		res = max(res, i-l)
		seenAt[ch-'a'] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

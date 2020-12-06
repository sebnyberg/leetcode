package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_challenge(t *testing.T) {
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
	if len(s) <= 1 {
		return len(s)
	}
	var (
		res      int  // length of longest unique substring
		startIdx int  // index of first unique character
		uniqLen  int  = 1
		i        int  // current index
		r        rune // current rune
	)
	// TODO: allocate beforehand
	seenAt := map[rune]int{}

	for i, r = range s {
		// If rune is not in map, add it and continue
		if _, exists := seenAt[r]; !exists {
			seenAt[r] = i
			continue
		}

		// Rune is in map (non-unique rune)
		// Store unique length if it was the greatest one
		uniqLen = i - startIdx
		if uniqLen > res {
			fmt.Println(i, startIdx, uniqLen)
			res = uniqLen
		}

		// Set new start index
		startIdx = seenAt[r] + 1

		// Update index
		seenAt[r] = i
	}

	// In case the last element was unique
	uniqLen = i - startIdx + 1
	fmt.Println(i, startIdx, uniqLen)
	if uniqLen > res {
		res = uniqLen
	}

	return res
}

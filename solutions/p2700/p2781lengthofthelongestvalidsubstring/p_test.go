package p2781lengthofthelongestvalidsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestValidSubstring(t *testing.T) {
	for i, tc := range []struct {
		word      string
		forbidden []string
		want      int
	}{
		{"cbaaaabc", []string{"aaa", "cb"}, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestValidSubstring(tc.word, tc.forbidden))
		})
	}
}

func longestValidSubstring(word string, forbidden []string) int {
	// There aren't that many ways for a string to be forbidden. We really only
	// need to check whether the final m characters are valid to keep adding
	// strings.
	m := make(map[string]bool)
	for _, w := range forbidden {
		m[w] = true
	}
	var l int
	var res int
	for r := range word {
		d := r - l + 1
		for k := 1; k <= min(10, d); k++ {
			if m[word[r-k+1:r+1]] {
				l = r - k + 2
				break
			}
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

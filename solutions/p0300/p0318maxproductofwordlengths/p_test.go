package p0318maxproductofwordlengths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProduct(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
		{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{[]string{"a", "aa", "aaa", "aaaa"}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, maxProduct(tc.words))
		})
	}
}

func maxProduct(words []string) int {
	// For each word, create an integer with bits per character
	// Then, bitwise AND with other words to check overlap
	// O(n^2), i.e. 1000000, should be OK
	n := len(words)
	wordNums := make([]int, n)
	for i, word := range words {
		var wordNum int
		for _, ch := range word {
			wordNum |= 1 << (ch - 'a')
		}
		wordNums[i] = wordNum
	}

	var maxResult int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if wordNums[i]&wordNums[j] == 0 {
				maxResult = max(maxResult, len(words[i])*len(words[j]))
			}
		}
	}
	return maxResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

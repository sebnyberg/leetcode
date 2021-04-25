package p1839longestsubstringovallvowelsinorder

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestBeautifulSubstring(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int
	}{
		{"aeiaaioaaaaeiiiiouuuooaauuaeiu", 13},
		{"aeeeiiiioooauuuaeiou", 5},
		{"a", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, longestBeautifulSubstring(tc.word))
		})
	}
}

func longestBeautifulSubstring(word string) int {
	// Intuition
	// If a string is sorted and includes all vowels, then it must start
	// with an 'a'
	// Thus, we can iterate over indices of 'a' in the input word,
	// adding letters until either a letter is smaller than it's previous,
	// or we've reached the end.

	res := 0
	var pos int
	idx := strings.IndexRune(word[pos:], 'a')
	if idx == -1 {
		return res
	}
	n := len(word)
	for {
		var i int
		nchars := 1
		for i = idx + 1; i < n && word[i] >= word[i-1]; i++ {
			if word[i] != word[i-1] {
				nchars++
			}
		}
		if nchars == 5 {
			res = max(res, i-idx)
		}
		pos = i
		nextIdx := strings.IndexRune(word[pos:], 'a')
		if nextIdx == -1 {
			return res
		}
		idx = pos + nextIdx
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

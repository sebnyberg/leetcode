package p1048longeststringchain

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestStrChain(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, longestStrChain(tc.words))
		})
	}
}

func longestStrChain(words []string) int {
	wordChains := make(map[string]int)
	maxLen := 0
	for _, word := range words {
		wordChains[word] = 1
		maxLen = max(maxLen, len(word))
	}
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })

	for _, word := range words {
		// Remove each letter from word and do a check of the word in wordChains
		for i := range word {
			k := word[:i] + word[i+1:]
			if _, exists := wordChains[k]; exists {
				wordChains[k] = max(wordChains[k], wordChains[word]+1)
			}
		}
	}

	maxVal := 0
	for _, n := range wordChains {
		maxVal = max(maxVal, n)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

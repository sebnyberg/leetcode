package p0140wordbreak2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordBreak(t *testing.T) {
	for _, tc := range []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cats and dog", "cat sand dog"}},
		{"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple"},
		},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.s, tc.wordDict), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, wordBreak(tc.s, tc.wordDict))
		})
	}
}

func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	if !canBreak(s, wordDict) {
		return []string{}
	}
	matchedPos := make([]bool, n+1)
	matchedPos[0] = true
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	results := make(map[int][]string)
	results[0] = []string{""}

	var start, pos int
	for start = 0; start < n; start++ {
		if !matchedPos[start] {
			continue
		}
		for pos = start + 1; pos <= n; pos++ {
			if words[s[start:pos]] && matchedPos[start] {
				matchedPos[pos] = true
				// For each word in the start position, add the matched word
				// word to that list and put it in the current position
				for _, result := range results[start] {
					results[pos] = append(results[pos], result+s[start:pos]+" ")
				}
			}
		}
	}
	for i := range results[n] {
		results[n][i] = results[n][i][:len(results[n][i])-1]
	}
	return results[n]
}

func canBreak(s string, wordDict []string) bool {
	covered := make([]bool, len(s)+1)
	covered[0] = true
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}

	var start, pos int
	for start = 0; start < len(s); start++ {
		for pos = start + 1; pos <= len(s); pos++ {
			covered[pos] = covered[pos] || covered[start] && words[s[start:pos]]
		}
	}
	return covered[len(s)]
}

package p0139wordbreak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordBreak(t *testing.T) {
	for _, tc := range []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.s, tc.wordDict), func(t *testing.T) {
			require.Equal(t, tc.want, wordBreak(tc.s, tc.wordDict))
		})
	}
}

func wordBreak(s string, wordDict []string) bool {
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

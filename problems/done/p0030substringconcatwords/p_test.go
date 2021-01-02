package p0030substringconcatwords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSubstring(t *testing.T) {
	for _, tc := range []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findSubstring(tc.s, tc.words))
		})
	}
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	nwords, wordLen := len(words), len(words[0])

	wordToIndex := make(map[string]int)
	wordCounts := make([]int, 0, len(words))
	i := 0
	for _, word := range words {
		wordIndex, exists := wordToIndex[word]
		if !exists {
			wordToIndex[word] = i
			wordCounts = append(wordCounts, 1)
			i++
			continue
		}
		wordCounts[wordIndex]++
	}

	wordMatchCount := make([]int, len(wordCounts))

	res := make([]int, 0)
	last := len(s) - (nwords * wordLen)
	for i := 0; i <= last; i++ {
		copy(wordMatchCount, wordCounts)
		for j := 0; j < nwords*wordLen; j += wordLen {
			substring := s[i+j : i+j+wordLen]
			if _, exists := wordToIndex[substring]; !exists {
				goto ContinueSearch
			}
			wordIndex := wordToIndex[substring]
			if wordMatchCount[wordIndex] == 0 {
				goto ContinueSearch
			}
			wordMatchCount[wordIndex]--
		}
		res = append(res, i)
	ContinueSearch:
	}

	return res
}

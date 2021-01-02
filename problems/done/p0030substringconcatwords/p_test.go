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

func Test_matchMap(t *testing.T) {
	for _, tc := range []struct {
		s       string
		wordLen int
		words   map[string]int
		want    bool
	}{
		// {"foobar", 3, map[string]int{"foo": 1, "bar": 1}, true},
		{"goodgoodbestword", 4, map[string]int{"word": 1, "good": 2, "best": 1}, true},
	} {
		t.Run(fmt.Sprintf("%v/%v/%+v", tc.s, tc.wordLen, tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, matchMap(tc.s, tc.wordLen, tc.words))
		})
	}
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	wordMap := make(map[string]int)
	wordLen := len(words[0])
	var wordsLen int
	for _, word := range words {
		wordMap[word]++
		wordsLen += len(word)
	}

	idx := 0

	res := make([]int, 0)
	for len(s[idx:]) >= wordsLen {
		if matchMap(s[idx:idx+wordsLen], wordLen, wordMap) {
			res = append(res, idx)
		}

		idx++
	}

	return res
}

func matchMap(s string, wordlen int, words map[string]int) bool {
	var idx int
	foundCount := make(map[string]int)
	for idx < len(s) {
		ss := s[idx : idx+wordlen]
		count, exists := words[ss]
		if !exists {
			return false
		}
		if foundCount[ss] >= count {
			return false
		}
		foundCount[ss]++
		idx += wordlen
	}
	return true
}

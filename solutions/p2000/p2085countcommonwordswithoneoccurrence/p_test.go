package p2085countcommonwordswithoneoccurrence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countWords(t *testing.T) {
	for _, tc := range []struct {
		words1 []string
		words2 []string
		want   int
	}{
		{[]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}, 2},
		{[]string{"b", "bb", "bbb"}, []string{"a", "aa", "aaa"}, 0},
		{[]string{"a", "ab"}, []string{"a", "a", "a", "ab"}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words1), func(t *testing.T) {
			require.Equal(t, tc.want, countWords(tc.words1, tc.words2))
		})
	}
}

func countWords(words1 []string, words2 []string) int {
	word1Count := make(map[string]int)
	for _, w := range words1 {
		word1Count[w]++
	}
	word2Count := make(map[string]int)
	for _, w := range words2 {
		word2Count[w]++
	}
	var res int
	for w, count := range word1Count {
		if count == 1 && word2Count[w] == 1 {
			res++
		}
	}
	return res
}

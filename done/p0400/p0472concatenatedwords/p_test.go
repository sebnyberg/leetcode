package p0472concatenatedwords

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  []string
	}{
		{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
		{[]string{"cat", "dog", "catdog"}, []string{"catdog"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findAllConcatenatedWordsInADict(tc.words))
		})
	}
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var res []string
	wordsMap := make(map[string]struct{})
	var empty [1001]bool
	empty[0] = true
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		dp := empty
		for i := 1; i <= len(w); i++ {
			for j := 0; j < i; j++ {
				if !dp[j] {
					continue
				}
				if _, exists := wordsMap[w[j:i]]; exists {
					dp[i] = true
					break
				}
			}
		}
		if dp[len(w)] {
			res = append(res, w)
		}
		wordsMap[w] = struct{}{}
	}
	return res
}

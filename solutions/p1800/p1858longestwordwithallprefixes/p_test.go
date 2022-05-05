package p1858longestwordwithallprefixes

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestWord(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{[]string{"k", "ki", "kir", "kira", "kiran"}, "kiran"},
		{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
		{[]string{"abc", "bc", "ab", "qwe"}, ""},
		{[]string{"abc", "ab", "a", "q", "qw", "qwe"}, "abc"},
		{[]string{"abc", "b", "ba", "ab", "a", "q", "qw", "qwe"}, "abc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, longestWord(tc.words))
		})
	}
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})
	var maxLen int
	var maxStr string
	prefixes := make(map[string]struct{}, len(words))
	for _, word := range words {
		if len(word) > 1 {
			if _, exists := prefixes[word[:len(word)-1]]; !exists {
				continue
			}
		}
		prefixes[word] = struct{}{}
		if l := len(word); l > maxLen {
			maxLen = l
			maxStr = word
		}
		continue
	}

	return maxStr
}

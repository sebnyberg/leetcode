package p0890findandreplacepattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAndReplacePattern(t *testing.T) {
	for _, tc := range []struct {
		words   []string
		pattern string
		want    []string
	}{
		{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb", []string{"mee", "aqq"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findAndReplacePattern(tc.words, tc.pattern))
		})
	}
}

func findAndReplacePattern(words []string, pattern string) []string {
	buf := make([]byte, len(pattern))
	normalize := func(s string) string {
		i := byte('a')
		var mapping [26]byte
		for j, ch := range s {
			if mapping[ch-'a'] == 0 {
				mapping[ch-'a'] = i
				i++
			}
			buf[j] = mapping[ch-'a']
		}
		return string(buf)
	}
	pattern = normalize(pattern)
	var res []string
	for _, w := range words {
		if normalize(w) == pattern {
			res = append(res, w)
		}
	}
	return res
}

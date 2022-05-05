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
	// re-format each word so that the first unique character becomes 'a',
	// second becomes 'b', etc
	// aqq => abb
	// cctqq => aabcc
	transPat := transform(pattern)
	res := make([]string, 0)
	for _, word := range words {
		if t := transform(word); t == transPat {
			res = append(res, word)
		}
	}
	return res
}

func transform(a string) string {
	var charPos [26]int
	res := make([]byte, len(a))
	pos := 1
	for i, ch := range a {
		if charPos[ch-'a'] == 0 {
			charPos[ch-'a'] = pos
			pos++
		}
		res[i] = byte(charPos[ch-'a'] - 1 + 'a')
	}
	return string(res)
}

package p2131longestpalindromebyconcatenatingtwoletterwords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"lc", "cl", "gg"}, 6},
		{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8},
		{[]string{"cc", "ll", "xx"}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, longestPalindrome(tc.words))
		})
	}
}

func longestPalindrome(words []string) int {
	pairCount := make(map[string]int)
	for _, word := range words {
		pairCount[word]++
	}
	var res int
	var hasMiddle bool
	for word, count := range pairCount {
		r := rev(word)
		if r == word {
			if count%2 == 1 {
				hasMiddle = true
			}
			res += (count / 2) * 4
		} else {
			n := pairCount[r]
			pairCount[r] -= min(n, count)
			pairCount[word] -= min(n, count)
			res += min(n, count) * 4
		}
	}
	if hasMiddle {
		res += 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

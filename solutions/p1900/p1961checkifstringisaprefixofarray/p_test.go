package p1961checkifstringisaprefixofarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPrefixString(t *testing.T) {
	for _, tc := range []struct {
		s     string
		words []string
		want  bool
	}{
		{"iloveleetcode", []string{"i", "love", "leetcode", "apples"}, true},
		{"iloveleetcode", []string{"apples", "i", "love", "leetcode"}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isPrefixString(tc.s, tc.words))
		})
	}
}

func isPrefixString(s string, words []string) bool {
	var count int
	for _, w := range words {
		if len(w) > len(s) || s[:len(w)] != w {
			break
		}
		count++
		s = s[len(w):]
	}
	return len(s) == 0
}

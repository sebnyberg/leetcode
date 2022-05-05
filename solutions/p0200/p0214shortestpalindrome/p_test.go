package p0214shortestpalindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"aabba", "abbaabba"},
		{"abb", "bbabb"},
		{"aacecaaa", "aaacecaaa"},
		{"abcd", "dcbabcd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, shortestPalindrome(tc.s))
		})
	}
}

func shortestPalindrome(s string) string {
	n := len(s)
	r := rev(s)
	res := kmp(s + "#" + r)
	return r[:n-res[n*2]] + s
}

func rev(s string) string {
	bs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func kmp(s string) []int {
	f := make([]int, len(s))
	f[0] = 0
	var j int
	for i := 1; i < len(s); i++ {
		j = f[i-1]
		// If j > 0 and s[i] == s[j], then
		// we can simply increment the previous prefix
		// match by one and continue.
		// If j == 0, then there is nothing left to do anyway
		for j > 0 && s[i] != s[j] {
			j = f[j-1]
		}
		if s[i] == s[j] {
			f[i] = j + 1
		}
	}
	return f
}

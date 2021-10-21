package p0214shortestpalindromerollinghash

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

const mod = 1_000_000_007

func shortestPalindrome(s string) string {
	pow := 1
	h1, h2 := 0, 0
	base := 29 // first prime above 26
	pos := -1
	for i := range s {
		h1 = (h1*base + int(s[i]-'a')) % mod
		h2 = (h2 + (int(s[i]-'a') * pow)) % mod
		if h1 == h2 {
			pos = i
		}
		pow = pow * base % mod
	}
	return rev(s[pos+1:]) + s
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

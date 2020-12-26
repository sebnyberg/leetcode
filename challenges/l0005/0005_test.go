package l0005_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_l0005(t *testing.T) {
	tcs := []struct {
		in   string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"ccc", "ccc"},
		{"bananas", "anana"},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, l0005longestPalindrome(tc.in))
		})
	}
}

func Benchmark_l0005(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l0005longestPalindrome("bananas")
	}
}

func l0005longestPalindrome(s string) string {
	// mid, left, right
	var m, l, r, nmax, maxL, maxR int
	for m < len(s) {
		if nmax/2 > len(s)-m {
			break
		}

		// Look for odd palindrome
		l, r = m, m
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > nmax {
			nmax = r - l - 1
			maxL, maxR = l, r
		}

		// Look for an even palindrome
		l, r = m-1, m
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > nmax {
			nmax = r - l - 1
			maxL, maxR = l, r
		}

		// move m one step to the right
		m++
	}
	return s[maxL+1 : maxR]
}

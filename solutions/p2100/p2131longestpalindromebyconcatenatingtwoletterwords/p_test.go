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
	m := make(map[[2]byte]int)
	for _, w := range words {
		m[[2]byte{w[0], w[1]}]++
	}
	var middle int
	var pairCount int
	for w, x := range m {
		if w[0] > w[1] {
			continue
		}
		if w[0] == w[1] {
			middle |= x & 1
		} else {
			ss := [2]byte{w[1], w[0]}
			if m[ss] < x {
				x = m[ss]
			}
			x *= 2
		}
		pairCount += x - x&1
	}
	return (pairCount + middle) * 2
}

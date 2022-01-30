package p0409longestpalindrome

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestPalindrome(tc.s))
		})
	}
}

func longestPalindrome(s string) int {
	var count [26 * 2]int
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			count[int(ch-'A')+26]++
		} else {
			count[int(ch-'a')]++
		}
	}
	var res, odd int
	for _, n := range count {
		odd |= n & 1
		res += (n / 2) * 2
	}
	return res + odd
}

package p0345reversevowelsofastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseVowels(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"aA", "Aa"},
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, reverseVowels(tc.s))
		})
	}
}

func reverseVowels(s string) string {
	var l, r int = 0, len(s) - 1
	var vowel [256]bool
	for _, ch := range []byte{
		'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U',
	} {
		vowel[ch] = true
	}
	bs := []byte(s)
	for l < r {
		switch {
		case !vowel[s[l]]:
			l++
		case !vowel[s[r]]:
			r--
		default:
			bs[l], bs[r] = bs[r], bs[l]
			l++
			r--
		}
	}
	return string(bs)
}

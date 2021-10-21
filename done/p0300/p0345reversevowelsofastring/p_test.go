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
	vowels := map[byte]bool{
		'a': true, 'A': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}
	bs := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for l <= r && !vowels[bs[l]] {
			l++
		}
		for l <= r && !vowels[bs[r]] {
			r--
		}
		if l < r {
			bs[l], bs[r] = bs[r], bs[l]
			l++
			r--
		}
	}
	return string(bs)
}

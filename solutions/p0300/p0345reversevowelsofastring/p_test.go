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
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, reverseVowels(tc.s))
		})
	}
}

var vowel = [256]uint8{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
}
var swap = [256][256]uint8{
	'a': vowel, 'e': vowel, 'i': vowel, 'o': vowel, 'u': vowel,
	'A': vowel, 'E': vowel, 'I': vowel, 'O': vowel, 'U': vowel,
}

func reverseVowels(s string) string {
	bs := []byte(s)
	var l, r int = 0, len(s) - 1
	for l < r {
		a, b := s[l], s[r]
		targetL := l + int(swap[a][b])*(r-l)
		targetR := r - int(swap[a][b])*(r-l)
		bs[targetL], bs[targetR] = bs[l], bs[r]
		moveR := ^(^vowel[a] & vowel[b]) & 1
		moveL := ^(vowel[a] & ^vowel[b]) & 1
		l += int(moveL)
		r -= int(moveR)
	}
	return string(bs)
}

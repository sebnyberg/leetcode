package p0408validwordabbreviation

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_validWordAbbreviation(t *testing.T) {
	for _, tc := range []struct {
		word string
		abbr string
		want bool
	}{
		{"internationalization", "i12iz4n", true},
		{"apple", "a2e", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, validWordAbbreviation(tc.word, tc.abbr))
		})
	}
}

func validWordAbbreviation(word string, abbr string) bool {
	var i, j int
	for i < len(abbr) && j < len(word) {
		if digit(abbr[i]) {
			if abbr[i] == '0' {
				return false
			}
			k := i
			for ; k < len(abbr) && digit(abbr[k]); k++ {
			}
			j += parseNum(abbr[i:k])
			i = k
			continue
		}
		if abbr[i] != word[j] {
			return false
		}
		i++
		j++
	}
	return i == len(abbr) && j == len(word)
}

func digit(b byte) bool {
	return unicode.Is(unicode.Digit, rune(b))
}

func parseNum(bs string) int {
	var res int
	for _, b := range bs {
		res *= 10
		res += int(b - '0')
	}
	return res
}

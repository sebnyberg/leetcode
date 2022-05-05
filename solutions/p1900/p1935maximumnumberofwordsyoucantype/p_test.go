package p1935maximumnumberofwordsyoucantype

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canBeTypedWords(t *testing.T) {
	for _, tc := range []struct {
		text          string
		brokenLetters string
		want          int
	}{
		{"hello world", "ad", 1},
		{"leet code", "lt", 1},
		{"leet code", "e", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.text), func(t *testing.T) {
			require.Equal(t, tc.want, canBeTypedWords(tc.text, tc.brokenLetters))
		})
	}
}

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	var broken [26]bool
	for _, lt := range brokenLetters {
		broken[lt-'a'] = true
	}
	var res int
	for _, word := range words {
		ok := true
		for _, ch := range word {
			if broken[ch-'a'] {
				ok = false
			}
		}
		if ok {
			res++
		}
	}
	return res
}

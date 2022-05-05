package p0316removeduplicateletters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicateLetters(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"cbacdcbc", "acdb"},
		{"bcabc", "abc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removeDuplicateLetters(tc.s))
		})
	}
}

func removeDuplicateLetters(s string) string {
	var letterCount [26]int
	for _, ch := range s {
		letterCount[ch-'a']++
	}
	letters := make([]rune, 0)
	n := 0
	var hasLetter [26]bool
	for _, ch := range s {
		letterCount[ch-'a']--
		if hasLetter[ch-'a'] {
			continue
		}
		for len(letters) > 0 && letters[n-1] > ch && letterCount[letters[n-1]-'a'] > 0 {
			hasLetter[letters[n-1]-'a'] = false
			letters = letters[:n-1] // pop
			n--
		}
		hasLetter[ch-'a'] = true
		n++
		letters = append(letters, ch)
	}

	return string(letters)
}

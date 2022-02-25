package p0520detectcapital

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_detectCapitalUse(t *testing.T) {
	for _, tc := range []struct {
		word string
		want bool
	}{
		{"Usa", true},
		{"USA", true},
		{"usa", true},
		{"FlaG", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, detectCapitalUse(tc.word))
		})
	}
}

func detectCapitalUse(word string) bool {
	if unicode.IsLower(rune(word[len(word)-1])) {
		// All must be lower
		for i := 1; i < len(word); i++ {
			if !unicode.IsLower(rune(word[i])) {
				return false
			}
		}
		return true
	}
	for i := range word {
		if !unicode.IsUpper(rune(word[i])) {
			return false
		}
	}
	return true
}

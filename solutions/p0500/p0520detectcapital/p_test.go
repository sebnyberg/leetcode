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
		{"FFFf", false},
		{"USA", true},
		{"FlaG", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, detectCapitalUse(tc.word))
		})
	}
}

func detectCapitalUse(word string) bool {
	isUpper := func(ch byte) bool {
		return unicode.Is(unicode.Upper, rune(ch))
	}
	isLower := func(ch byte) bool {
		return unicode.Is(unicode.Lower, rune(ch))
	}

	for i := 0; i < len(word)-1; i++ {
		if i > 0 && isUpper(word[i]) && isLower(word[i+1]) ||
			isLower(word[i]) && isUpper(word[i+1]) {
			return false
		}
	}
	return true
}

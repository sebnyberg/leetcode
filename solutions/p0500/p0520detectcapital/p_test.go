package p0520detectcapital

import (
	"fmt"
	"strings"
	"testing"

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
	lo := strings.ToLower(word)
	up := strings.ToUpper(word)
	return up == word || word[1:] == lo[1:]
}

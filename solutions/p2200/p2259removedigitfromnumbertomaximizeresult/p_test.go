package p2259removedigitfromnumbertomaximizeresult

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDigit(t *testing.T) {
	for _, tc := range []struct {
		number string
		digit  byte
		want   string
	}{
		{"123", '3', "12"},
		{"1231", '1', "231"},
		{"551", '5', "51"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.number), func(t *testing.T) {
			require.Equal(t, tc.want, removeDigit(tc.number, tc.digit))
		})
	}
}

func removeDigit(number string, digit byte) string {
	// If we can remove a digit so that it increases in value compared to the next
	// number, then it is optimal.
	// If this is not the case, then we wish to remove the last occurrence.
	pos := -1
	for i, ch := range number {
		if byte(ch) == digit {
			if i < len(number)-1 && byte(number[i+1]) > byte(ch) {
				// Done
				return number[:i] + number[i+1:]
			}
			pos = i
		}
	}
	return number[:pos] + number[pos+1:]
}

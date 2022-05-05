package p0151reversewords

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"  a word", "word a"},
		{"  a word  ", "word a"},
		{"a word  ", "word a"},
		{"a  word  ", "word a"},
		{"a  word", "word a"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, reverseWords(tc.s))
		})
	}
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	for l, r := 0, len(words)-1; l < r; l, r = l+1, r-1 {
		words[l], words[r] = words[r], words[l]
	}
	return strings.Join(words, " ")
}

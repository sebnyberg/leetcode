package p0686repeatedstringmatch

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_repeatedStringMatch(t *testing.T) {
	for i, tc := range []struct {
		a    string
		b    string
		want int
	}{
		{"abcd", "abc", 1},
		{"abcd", "abcda", 2},
		{"abcd", "cdabcdab", 3},
		{"a", "aa", 2},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, repeatedStringMatch(tc.a, tc.b))
		})
	}
}

func repeatedStringMatch(a string, b string) int {
	k := max(1, len(b)/len(a))
	c := strings.Repeat(a, k*3)
	for i := 0; i+len(b) < len(c); i++ {
		if c[i:i+len(b)] == b {
			pos := i + len(b) - 1
			return (pos / len(a)) + 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

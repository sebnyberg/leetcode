package p0290wordpattern

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordPattern(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	} {
		t.Run(fmt.Sprintf("%+v,%+v", tc.pattern, tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, wordPattern(tc.pattern, tc.s))
		})
	}
}

func wordPattern(pattern string, s string) bool {
	seen := make(map[byte]string)
	words := make(map[string]bool)
	parts := strings.Split(s, " ")
	if len(parts) != len(pattern) {
		return false
	}
	for i, w := range strings.Split(s, " ") {
		k := pattern[i]
		if v, exists := seen[k]; !exists {
			if words[w] {
				return false
			}
			seen[k] = w
			words[w] = true
			continue
		} else {
			if v != w {
				return false
			}
		}
	}
	return true
}

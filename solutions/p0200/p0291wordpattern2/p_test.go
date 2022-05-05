package p0291wordpattern2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordPattern(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		s       string
		want    bool
	}{
		{"ab", "aa", false},
		{"aaaa", "asdasdasdasd", true},
		{"abab", "asdasdasdasd", true},
		{"aabb", "xyzabcxzyabc", false},
		{"abab", "redblueredblue", true},
	} {
		t.Run(fmt.Sprintf("%+v,%+v", tc.pattern, tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, wordPatternMatch(tc.pattern, tc.s))
		})
	}
}

func wordPatternMatch(pattern string, s string) bool {
	return helper(pattern, 0, s, 0, map[byte]string{}, map[string]struct{}{})
}

func helper(pattern string, patPos int, s string, sPos int, patString map[byte]string, seen map[string]struct{}) bool {
	if patPos == len(pattern) {
		return sPos == len(s)
	}
	p := pattern[patPos]
	if _, exists := patString[p]; !exists {
		for i := 1; sPos+i <= len(s); i++ {
			k := s[sPos : sPos+i]
			if _, exists := seen[k]; exists {
				continue
			}
			patString[p] = k
			seen[k] = struct{}{}
			if helper(pattern, patPos+1, s, sPos+i, patString, seen) {
				return true
			}
			delete(seen, k)
			delete(patString, p)
		}
		return false
	}
	// Pattern exist, match until out of bounds
	patStr := patString[p]
	if sPos+len(patStr) > len(s) {
		return false
	}
	for i := 0; i < len(patStr); i++ {
		if s[sPos+i] != patStr[i] {
			return false
		}
	}
	return helper(pattern, patPos+1, s, sPos+len(patStr), patString, seen)
}

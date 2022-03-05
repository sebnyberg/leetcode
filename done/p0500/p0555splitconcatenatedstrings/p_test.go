package p0555splitconcatenatedstrings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitLoopedString(t *testing.T) {
	for _, tc := range []struct {
		strs []string
		want string
	}{
		{[]string{"abc", "xyz"}, "zyxcba"},
		{[]string{"abc"}, "cba"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, splitLoopedString(tc.strs))
		})
	}
}

func splitLoopedString(strs []string) string {
	// Find maximum character, e.g. 'z'
	var maxChar rune
	for _, s := range strs {
		for _, ch := range s {
			if ch > maxChar {
				maxChar = ch
			}
		}
	}

	// Make all strs flipped in lexicographically greatest order
	for i := range strs {
		r := rev(strs[i])
		if r > strs[i] {
			strs[i] = r
		}
	}

	// Try each possible starting position
	var cand string
	for i := range strs {
		s := strs[i]
		r := rev(s)
		pre := strings.Join(strs[:i], "")
		post := strings.Join(strs[i+1:], "")
		for j := 0; j < len(s); j++ {
			if s[j] != byte(maxChar) {
				continue
			}
			// If this index is the starting position of the solution, then it could
			// go in either directions.
			forwardCand := s[j:] + post + pre + s[:j]
			if forwardCand > cand {
				cand = forwardCand
			}
			// Try also the other direction
			ri := len(s) - j - 1
			revCand := r[ri:] + post + pre + r[:ri]
			if revCand > cand {
				cand = revCand
			}
		}
	}
	return cand
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

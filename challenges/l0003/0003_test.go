package l0003_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_l0003(t *testing.T) {
	tcs := []struct {
		in      string
		want    int
		wantStr string
	}{
		{"aab", 2, "ab"},
		{"dvdf", 3, "vdf"},
		{"abcabcbb", 3, "abc"},
		{"bbbbb", 1, "b"},
		{"pwwkew", 3, "wke"},
		{"", 0, ""},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, l0003_lengthOfLongestSubstring(tc.in))
		})
	}
}

func l0003_lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var (
		nuniq  int  = 1
		i      int  // current index
		r      rune // current rune
		gcFrom int  // previous end index
	)
	// TODO: allocate beforehand
	seenAt := map[rune]int{}

	for i, r = range s {
		// If rune is not in map, add it and continue
		if _, exists := seenAt[r]; !exists {
			seenAt[r] = i
			continue
		}

		if len(seenAt) > nuniq {
			nuniq = len(seenAt)
		}

		// Clean from previous gc until (but not including) s[seenAt[r]]
		for _, k := range s[gcFrom:seenAt[r]] {
			delete(seenAt, k)
		}
		gcFrom = seenAt[r] + 1

		seenAt[r] = i
	}

	// In case the last element was unique
	if len(seenAt) > nuniq {
		nuniq = len(seenAt)
	}

	return nuniq
}

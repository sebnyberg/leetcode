package p2730findthelongestsemirepetitivesubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSemiRepetitiveSubstring(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"52233", 4},
		{"11111111", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestSemiRepetitiveSubstring(tc.s))
		})
	}
}

func longestSemiRepetitiveSubstring(s string) int {
	var l int
	var b bool
	var res int
	for i := 0; i < len(s); i++ {
		res = max(res, i-l+1)
		if i < len(s)-1 && s[i+1] == s[i] {
			if !b {
				b = true
			} else {
				for s[l] != s[l+1] {
					l++
				}
				l++
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

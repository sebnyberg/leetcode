package p2414lengthofthelongestalphabeticalcontinuoussubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestContinuousSubstring(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"abcde", 5},
		{"abacaba", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestContinuousSubstring(tc.s))
		})
	}
}

func longestContinuousSubstring(s string) int {
	res := 1
	m := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			m++
			res = max(res, m)
			continue
		}
		m = 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

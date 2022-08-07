package p2370longestidealsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestIdealString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"azaza", 25, 5},
		{"acfgbd", 2, 4},
		{"abcd", 3, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestIdealString(tc.s, tc.k))
		})
	}
}

func longestIdealString(s string, k int) int {
	// curr contains the max substring len ending in a certain character.
	var curr [26]int
	for _, ch := range s {
		next := curr
		a := int(ch - 'a')
		// Append to characters above/below
		for b := max(0, a-k); b <= min(25, a+k); b++ {
			next[a] = max(next[a], 1+curr[b])
		}
		next[a] = max(next[a], 1) // can always form length 1
		curr = next
	}
	var res int
	for _, a := range curr {
		res = max(res, a)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

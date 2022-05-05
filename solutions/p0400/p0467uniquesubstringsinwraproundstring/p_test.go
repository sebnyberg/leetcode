package p0467uniquesubstringsinwraproundstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSubstringInWraproundString(t *testing.T) {
	for _, tc := range []struct {
		p    string
		want int
	}{
		{"yhxtdobyly", 8},
		{"a", 1},
		{"cac", 2},
		{"zab", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.p), func(t *testing.T) {
			require.Equal(t, tc.want, findSubstringInWraproundString(tc.p))
		})
	}
}

func findSubstringInWraproundString(p string) int {
	// Any string which is in cyclical alphabetical order is OK.

	// seq[char] stores the longest length sequence ending in char
	var seq [26]int

	var runLen int
	for i, ch := range p {
		k := ch - 'a'
		prev := byte(((k-1)+26)%26 + 'a')
		if i > 0 && p[i-1] != prev {
			runLen = 1
		} else {
			runLen++
		}
		seq[k] = max(seq[k], runLen)
	}
	var res int
	for _, n := range seq {
		res += n
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

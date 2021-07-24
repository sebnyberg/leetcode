package p1930uniquelength3palindromicsubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPalindromicSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"bbcbaba", 4},
		{"aabca", 3},
		{"adc", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countPalindromicSubsequence(tc.s))
		})
	}
}

func countPalindromicSubsequence(s string) int {
	var singleExists [26]bool
	var doubleExists [26][26]bool
	var tripleExists [26][26][26]bool

	for _, a := range s {
		for b, exists := range doubleExists[a-'a'] {
			if exists {
				tripleExists[a-'a'][b][a-'a'] = true
			}
		}

		for c, exists := range singleExists {
			if exists {
				doubleExists[c][a-'a'] = true
			}
		}

		singleExists[a-'a'] = true
	}

	var res int
	for i := range tripleExists {
		for j := range tripleExists[i] {
			for _, exists := range tripleExists[i][j] {
				if exists {
					res++
				}
			}
		}
	}
	return res
}

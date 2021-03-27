package p0647palindromicsubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countSubstrings(tc.s))
		})
	}
}

func countSubstrings(s string) int {
	var count int
	for i := range s {
		count += checkPalindrome(s, i, i)
		count += checkPalindrome(s, i, i+1)
	}
	return count
}

func checkPalindrome(s string, l, r int) int {
	var count int
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		count++
	}
	return count
}

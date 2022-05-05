package p2108findfirstpalindromicstringinthearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstPalindrome(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{[]string{"abc", "car", "ada", "racecar", "cool"}, "ada"},
		{[]string{"notapalindrome", "racecar"}, "racecar"},
		{[]string{"def", "ghi"}, ""},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, firstPalindrome(tc.words))
		})
	}
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""
}

package p1328breakpalindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_breakPalindrome(t *testing.T) {
	for _, tc := range []struct {
		palindrome string
		want       string
	}{
		{"abccba", "aaccba"},
		{"a", ""},
		{"aa", "ab"},
		{"aba", "abb"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.palindrome), func(t *testing.T) {
			require.Equal(t, tc.want, breakPalindrome(tc.palindrome))
		})
	}
}

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}
	isOdd := n%2 == 1
	for i, ch := range palindrome {
		if isOdd && i == n/2 {
			continue
		}
		if ch != 'a' {
			return palindrome[:i] + "a" + palindrome[i+1:]
		}
	}
	return palindrome[:n-1] + "b"
}

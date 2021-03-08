package p1332removepalindromicsubseq

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removePalindromeSub(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"ababa", 1},
		{"abb", 2},
		{"baabb", 2},
		{"", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, removePalindromeSub(tc.s))
		})
	}
}

func removePalindromeSub(s string) (res int) {
	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	if strings.ContainsRune(s, 'a') {
		res++
	}
	if strings.ContainsRune(s, 'b') {
		res++
	}
	return res
}

func isPalindrome(s string) bool {
	n := len(s)
	mid := n / 2
	l, r := mid-1, mid+1
	if n%2 == 0 {
		r = mid
	}
	for ; l >= 0 && r < n; l, r = l-1, r+1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

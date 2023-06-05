package p1745palindromepartitioningiv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPartitioning(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want bool
	}{
		{"abcbdd", true},
		{"bcbddxy", false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, checkPartitioning(tc.s))
		})
	}
}

func checkPartitioning(s string) bool {
	// If we split the string into three valid palindromes, one must be at the
	// center. Figuring out whether s[:l] is a palindrome is relatively easy.
	// Same with right side. If we memoize both (O(n)) we can simply test every
	// possible middle string against its sides to see if the three are valid.
	//
	buf := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	rev := string(buf)
	n := len(s)

	// isPalindrome(i, j) returns whether s[i:j] is a palindrome
	isPalindrome := func(i, j int) bool {
		m := j - i
		k := m / 2
		l := s[i : i+k]
		r := rev[n-j : n-j+k]
		return l == r
	}

	left := make([]bool, n)
	left[1] = true
	for i := 2; i < n; i++ {
		left[i] = isPalindrome(0, i)
	}

	right := make([]bool, n)
	right[n-1] = true
	for i := n - 2; i >= 2; i-- {
		right[i] = isPalindrome(i, n)
	}

	for l := 1; l < n-1; l++ {
		for r := n - 1; r > l; r-- {
			if left[l] && right[r] && isPalindrome(l, r) {
				return true
			}
		}
	}

	return false
}

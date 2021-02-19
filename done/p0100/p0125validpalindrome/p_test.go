package p0125validpalindrome

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"0P", false},
		{" ", true},
		{"abb", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isPalindrome(tc.s))
		})
	}
}

func isPalindrome(s string) bool {
	valid := "0123456789abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch -= 'A' - 'a'
		}
		if !strings.ContainsRune(valid, ch) {
			continue
		}
		sb.WriteRune(ch)
	}
	stripped := sb.String()
	n := len(stripped)
	if n <= 1 {
		return true
	}
	m := (n - 1) / 2
	for i, j := m, m+(1-n%2); i >= 0; i, j = i-1, j+1 {
		if stripped[i] != stripped[j] {
			return false
		}
	}
	return true
}

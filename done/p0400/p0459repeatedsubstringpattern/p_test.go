package p0459repeatedsubstringpattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, repeatedSubstringPattern(tc.s))
		})
	}
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for k := n / 2; k >= 1; k-- {
		if n%k != 0 {
			continue
		}
		var i int
		for i = k; i < n && s[i:i+k] == s[:k]; i += k {
		}
		if i == n {
			return true
		}
	}
	return false
}

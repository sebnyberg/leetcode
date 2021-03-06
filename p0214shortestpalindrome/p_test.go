package p0214shortestpalindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		// {"aacecaaa", "aaacecaaa"},
		// {"abcd", "dcbabcd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, shortestPalindrome(tc.s))
		})
	}
}

func shortestPalindrome(s string) string {
	return ""
	// for i := 0; i < len(s); i++ {
	// 	for j := 0; j < len(s); j++ {

	// 	}
	// }
}

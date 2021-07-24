package p1392longesthappyprefix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPrefix(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"level", "l"},
		{"ababab", "abab"},
		{"leetcodeleet", "leet"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestPrefix(tc.s))
		})
	}
}

func longestPrefix(s string) string {
	n := len(s)
	mod := 1_000_000_007
	// mod := (1 << 31) - 1
	// mod := 16_777_619_341_237
	base := 29
	h1, h2 := 0, 0
	pow := 1
	var maxRes string
	for i := 0; i < n-1; i++ {
		h1 = (h1*base + int(s[i]-'a')) % mod
		h2 = (int(s[n-1-i]-'a')*pow + h2) % mod
		pow = pow * base % mod
		if h1 == h2 {
			maxRes = s[:i+1]
		}
	}
	return maxRes
}

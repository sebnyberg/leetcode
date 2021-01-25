package p0010regexpmatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isMatch(t *testing.T) {
	tcs := []struct {
		s    string
		p    string
		want bool
	}{
		{"a", "ab*a", false},
		{"a", "ab*", true},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"aaa", "a*a", true},
		{"ab", ".*", true},
		{"ab", ".*c", false},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"mississippi", "mis*is*ip*.", true},
		{"aaa", "ab*ac*a", true},
		{"aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s", true},
		{"caccccaccbabbcb", "c*c*b*a*.*c*.a*a*a*", true},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.p), func(t *testing.T) {
			require.Equal(t, tc.want, isMatch(tc.s, tc.p))
		})
	}
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	match := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (match && isMatch(s[1:], p))
	}
	return match && isMatch(s[1:], p[1:])
}

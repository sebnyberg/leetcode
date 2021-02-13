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
		// {"a", "ab*a", false},
		// {"a", "ab*", true},
		// {"aa", "a", false},
		// {"aa", "a*", true},
		// {"aaa", "a*a", true},
		{"ab", ".*", true},
		// {"ab", ".*c", false},
		// {"aab", "c*a*b", true},
		// {"mississippi", "mis*is*p*.", false},
		// {"mississippi", "mis*is*ip*.", true},
		// {"aaa", "ab*ac*a", true},
		// {"aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s", true},
		// {"caccccaccbabbcb", "c*c*b*a*.*c*.a*a*a*", true},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.p), func(t *testing.T) {
			require.Equal(t, tc.want, isMatch(tc.s, tc.p))
		})
	}
}

func isMatch(s string, p string) bool {
	np := len(p)
	ns := len(s)
	dp := make([][]bool, ns+1)
	for i := range dp {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for i := 2; i <= np; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= ns; i++ {
		for j := 1; j <= np; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if j > 1 {
					dp[i][j] = dp[i][j-2]
					if p[j-2] == s[i-1] || p[j-2] == '.' {
						dp[i][j] = dp[i][j] || dp[i-1][j-2] || dp[i-1][j]
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

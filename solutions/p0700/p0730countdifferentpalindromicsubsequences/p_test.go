package p0730countdifferentpalindromicsubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPalindromicSubsequences(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"daaaabaadcccbcbbaabdbccbbabccacdacdddacbadbadcadcb", 99418},
		{"ddac", 4},
		{"bccb", 6},
		{"abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba", 104860361},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countPalindromicSubsequences(tc.s))
		})
	}
}

const mod = 1e9 + 7

func countPalindromicSubsequences(s string) int {
	var dp [4][][]int
	n := len(s)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var d func(start, end int, alpha byte) int
	d = func(start, end int, alpha byte) int {
		if start > end {
			return 0
		}
		if start == end {
			if s[start] == alpha {
				return 1
			}
			return 0
		}
		if dp[alpha-'a'][start][end] != -1 {
			return dp[alpha-'a'][start][end]
		}

		var wtf int
		if s[start] == s[end] && s[start] == alpha {
			wtf = 2
			for b := byte('a'); b <= 'd'; b++ {
				wtf = (wtf + d(start+1, end-1, b)) % mod
			}
		} else {
			wtf = (wtf + d(start, end-1, alpha)) % mod
			wtf = (wtf + d(start+1, end, alpha)) % mod
			wtf = (wtf - d(start+1, end-1, alpha)) % mod
			if wtf < 0 {
				wtf += mod
			}
		}
		dp[alpha-'a'][start][end] = wtf
		return wtf
	}

	var res int
	for a := byte('a'); a <= 'd'; a++ {
		res = (res + d(0, len(s)-1, a)) % mod
	}
	return res
}

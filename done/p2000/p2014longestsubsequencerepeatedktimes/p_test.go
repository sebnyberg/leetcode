package p2014longestsubsequencerepeatedktimes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSubsequenceRepeatedK(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"letsleetcode", 2, "let"},
		{"bb", 2, "b"},
		{"ab", 2, ""},
		{"bbabbabbbbabaababab", 3, "bbbb"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestSubsequenceRepeatedK(tc.s, tc.k))
		})
	}
}

func longestSubsequenceRepeatedK(s string, k int) string {
	n := 26
	res := ""
	cur := make([]string, 0)
	next := make([]string, 0)
	cur = append(cur, "")
	for len(cur) > 0 {
		next = next[:0]
		for _, curS := range cur {
			for i := 0; i < n; i++ {
				cand := curS + string(rune('a'+i))
				if isSub(s, cand, k) {
					res = cand
					next = append(next, cand)
				}
			}
		}
		next, cur = cur, next
	}
	return res
}

func TestIsSub(t *testing.T) {
	res := isSub("letsleetcode", "le", 2)
	require.True(t, res)
}

func isSub(s, sub string, k int) bool {
	var j, repeat int
	for i := 0; i < len(s); i++ {
		if s[i] == sub[j] {
			j++
			if j == len(sub) {
				repeat++
				if repeat == k {
					return true
				}
				j = 0
			}
		}
	}
	return false
}

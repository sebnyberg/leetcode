package p0438findallanagramsinstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAnagrams(t *testing.T) {
	for _, tc := range []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findAnagrams(tc.s, tc.p))
		})
	}
}

func findAnagrams(s string, p string) []int {
	var pattern [26]int16
	for _, ch := range p {
		pattern[ch-'a']++
	}
	n := len(p)
	var res []int
	var count [26]int16
	for i := range s {
		count[s[i]-'a']++
		if i >= n {
			count[s[i-n]-'a']--
		}
		if count == pattern {
			res = append(res, i-n+1)
		}
	}
	return res
}

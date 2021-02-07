package p0076minwindowsubst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minWindow(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want string
	}{
		{"cabefgecdaecf", "cae", "aec"},
		{"ab", "b", "b"},
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"bba", "ab", "ba"},
		{"a", "a", "a"},
		{"aa", "a", "a"},
		{"a", "aa", ""},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.s, tc.t), func(t *testing.T) {
			require.Equal(t, tc.want, minWindow(tc.s, tc.t))
		})
	}
}

func minWindow(s string, t string) (res string) {
	n := len(s)
	need := make(map[byte]int, len(t))
	for i := range t {
		need[t[i]]++
	}
	missing := len(t)
	l := 0
	for r := 0; r < n; r++ {
		if r == n {
			return res
		}
		if need[s[r]] > 0 {
			missing--
		}
		need[s[r]]--
		if missing == 0 {
			for ; l < r && need[s[l]] < 0; l++ {
				need[s[l]]++
			}
			need[s[l]]++
			missing++
			if res == "" || r-l+1 < len(res) {
				res = s[l : r+1]
			}
			l++
		}
	}

	return res
}

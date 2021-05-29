package p1876substringsofsizethreewithdistinctchars

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countGoodSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"xyzzaz", 1},
		{"aababcabc", 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countGoodSubstrings(tc.s))
		})
	}
}

func countGoodSubstrings(s string) int {
	var count [26]int
	n := len(s)
	if n < 3 {
		return 0
	}
	for i := 0; i < 3; i++ {
		count[s[i]-'a']++
	}
	var res int
	for i := 3; i <= n; i++ {
		ok := true
		for _, c := range count {
			if c > 1 {
				ok = false
			}
		}
		if ok {
			res++
		}
		if i == n {
			break
		}
		count[s[i-3]-'a']--
		count[s[i]-'a']++
	}
	return res
}

package p2223sumofscoresofbuiltstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumScores(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int64
	}{
		{"babab", 9},
		{"azbazbzaz", 14},
		// {"abcbabcb", 14},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, sumScores(tc.s))
		})
	}
}

func sumScores(s string) int64 {
	ans := 0
	n := len(s)
	z := make([]int, n)
	var l, r int
	for i := 1; i < n; i++ {
		k := max(0, min(z[i-l], r-i+1))
		for i+k < n && s[k] == s[i+k] {
			l, r = i, i+k
			k++
		}
		z[i] = k
	}
	z[0] = n

	for _, v := range z {
		ans += v
	}

	return int64(ans)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

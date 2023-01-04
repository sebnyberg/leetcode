package p1092shortestcommonsupersequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestCommonSupersequence(t *testing.T) {
	for i, tc := range []struct {
		str1 string
		str2 string
		want string
	}{
		{"abac", "cab", "cabac"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, shortestCommonSupersequence(tc.str1, tc.str2))
		})
	}
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	m := len(str1)
	n := len(str2)
	curr := make([]string, n+1)
	for j := range curr {
		curr[j] = str2[:j]
	}
	next := make([]string, n+1)
	for i := 1; i <= m; i++ {
		next = next[:1]
		next[0] = str1[:i]
		for j := 1; j <= n; j++ {
			down := curr[j] + string(str1[i-1])
			right := next[j-1] + string(str2[j-1])
			diag := curr[j-1] + string(str1[i-1])
			if str1[i-1] != str2[j-1] {
				diag += string(str2[j-1])
			}
			res := down
			if len(right) < len(res) {
				res = right
			}
			if len(diag) < len(res) {
				res = diag
			}
			next = append(next, res)
		}
		next, curr = curr, next
	}
	return curr[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

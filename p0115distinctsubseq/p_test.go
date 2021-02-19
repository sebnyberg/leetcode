package p0115distinctsubseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDistinct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"acdabefbc", "ab", 4},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.t), func(t *testing.T) {
			require.Equal(t, tc.want, numDistinct(tc.s, tc.t))
		})
	}
}

func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i, tch := range t {
		for j, sch := range s {
			if sch == tch {
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}

	return dp[m][n]
}

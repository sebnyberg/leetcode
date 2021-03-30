package p0279perfectsquares

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSquares(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{12, 3},
		{13, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, numSquares(tc.n))
		})
	}
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			dp[i+j*j] = min(dp[i+j*j], dp[i]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

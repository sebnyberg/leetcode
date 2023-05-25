package p0837new21game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_new21Game(t *testing.T) {
	for i, tc := range []struct {
		n      int
		k      int
		maxPts int
		want   float64
	}{
		{10, 1, 10, 1.0},
		{6, 1, 10, 0.6},
		{21, 17, 10, 0.73278},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.InEpsilon(t, tc.want, new21Game(tc.n, tc.k, tc.maxPts), eps)
		})
	}
}

const eps = 1e-5

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k+maxPts {
		return 1
	}
	dp := make([]float64, n+1)
	dp[0] = 1
	windowSum := 1.0
	var res float64
	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			res += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}
	return res
}

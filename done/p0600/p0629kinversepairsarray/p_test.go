package p0629kinversepairsarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kInversePairs(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{3, 0, 1},
		{3, 1, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, kInversePairs(tc.n, tc.k))
		})
	}
}

const mod = 1_000_000_007

func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}

	dp := make([]int, k+1)
	tmp := make([]int, k+1)
	for i := 1; i <= n; i++ {
		tmp[0] = 1
		for j := 1; j <= k; j++ {
			x := 0
			if j >= i {
				x = dp[j-i]
			}
			tmp[j] = (tmp[j-1] + dp[j] - x) % mod
		}
		dp, tmp = tmp, dp
	}
	return (dp[k] + mod - dp[k-1]) % mod
}

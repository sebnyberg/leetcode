package p1884eggdropwith2eggsandnfloors

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoEggDrop(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 2},
		{100, 14},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, twoEggDrop(tc.n))
		})
	}
}

func twoEggDrop(n int) int {
	// dp[eggs][floor]
	var dp [2][1001]int

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for floor := 1; floor <= n; floor++ {
		dp[1][floor] = math.MaxInt32
		for k := 1; k <= floor; k++ {
			c := 1 + max(dp[0][k-1], dp[1][floor-k])
			dp[1][floor] = min(dp[1][floor], c)
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

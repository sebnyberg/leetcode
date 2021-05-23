package p1866numberofwaystorearrangestickswithksticksvisible

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rearrangeSticks(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{3, 2, 3},
		{5, 5, 1},
		{20, 11, 647427950},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.n, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, rearrangeSticks(tc.n, tc.k))
		})
	}
}

const mod = 1000000007

var dp [1001][1001]int

func rearrangeSticks(n int, k int) int {
	if k == 0 {
		return 0
	}
	if n == k {
		return 1
	}
	if dp[n][k] == 0 {
		// Longest stick could be used in the last position
		dp[n][k] = rearrangeSticks(n-1, k-1)

		// Or, place any of the short sticks and expect the rest to be covered
		dp[n][k] += rearrangeSticks(n-1, k) * (n - 1)
		dp[n][k] %= mod
	}

	return dp[n][k]
}

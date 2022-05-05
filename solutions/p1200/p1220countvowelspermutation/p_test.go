package p1220countvowelspermutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowelPermutation(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countVowelPermutation(tc.n))
		})
	}
}

const mod = 1_000_000_007

func countVowelPermutation(n int) int {
	const (
		a = 0
		e = 1
		i = 2
		o = 3
		u = 4
	)

	var dp [2][5]int
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for it := 1; it < n; it++ {
		prev := (it - 1) % 2
		idx := it % 2
		dp[idx][a] = dp[prev][e] + dp[prev][i] + dp[prev][u]
		dp[idx][e] = dp[prev][a] + dp[prev][i]
		dp[idx][i] = dp[prev][e] + dp[prev][o]
		dp[idx][o] = dp[prev][i]
		dp[idx][u] = dp[prev][i] + dp[prev][o]
		if it%10 == 0 {
			for i := range dp[idx] {
				dp[idx][i] %= mod
			}
		}
	}

	var res int
	for _, val := range dp[(n-1)%2] {
		res += val
	}

	return res % mod
}

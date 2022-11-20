package p2478numberofbeautifulpartitions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautifulPartitions(t *testing.T) {
	for i, tc := range []struct {
		s         string
		k         int
		minLength int
		want      int
	}{
		{"23542185131", 3, 2, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, beautifulPartitions(tc.s, tc.k, tc.minLength))
		})
	}
}

const mod = 1e9 + 7

func beautifulPartitions(s string, k int, minLength int) int {
	// Traditional dp problem. Memoize the number of ways to form k partitions
	// with a given section of the string.
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	res := dfs(dp, s, k, minLength, 0)
	return res % mod
}

var prime = [10]bool{
	2: true,
	3: true,
	5: true,
	7: true,
}

func dfs(dp [][]int, s string, k, minLength, i int) int {
	if k == 0 {
		if i == len(s) {
			return 1
		}
		return 0
	}
	if i > len(s) || len(s)-i < minLength || !prime[s[i]-'0'] {
		return 0
	}
	if dp[k][i] != -1 {
		return dp[k][i]
	}
	var res int
	for j := i + minLength - 1; j < len(s); j++ {
		if prime[s[j]-'0'] {
			continue
		}
		res += dfs(dp, s, k-1, minLength, j+1)
	}
	dp[k][i] = res % mod
	return dp[k][i]
}

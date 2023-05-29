package p1563stonegamev

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stoneGameV(t *testing.T) {
	for i, tc := range []struct {
		stoneValue []int
		want       int
	}{
		{[]int{6, 2, 3, 4, 5, 5}, 18},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, stoneGameV(tc.stoneValue))
		})
	}
}

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	presum := make([]int, n+1)
	for i := range stoneValue {
		presum[i+1] = presum[i] + stoneValue[i]
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	res := dfs(dp, presum, 0, n)
	return res
}

func dfs(dp [][]int, presum []int, i, j int) int {
	if dp[i][j] != math.MinInt32 {
		return dp[i][j]
	}
	if j-i == 1 {
		return 0
	}
	var res int
	for k := i + 1; k < j; k++ {
		left := presum[k] - presum[i]
		right := presum[j] - presum[k]
		if right >= left {
			// keep right
			res = max(res, left+dfs(dp, presum, i, k))
		}
		if left >= right {
			// keep left
			res = max(res, right+dfs(dp, presum, k, j))
		}
	}
	dp[i][j] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

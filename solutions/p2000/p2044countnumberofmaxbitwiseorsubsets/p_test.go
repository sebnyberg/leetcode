package p2044countnumberofmaxbitwiseorsubsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countMaxOrSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 1}, 2},
		{[]int{2, 2, 2}, 7},
		{[]int{3, 2, 1, 5}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countMaxOrSubsets(tc.nums))
		})
	}
}

func countMaxOrSubsets(nums []int) int {
	var maxOR int
	for _, n := range nums {
		maxOR |= n
	}
	return dfs(nums, 0, len(nums), 0, maxOR)
}

func dfs(nums []int, i, n, cur, want int) int {
	if i == n {
		if cur == want {
			return 1
		}
		return 0
	}
	return dfs(nums, i+1, n, cur|nums[i], want) + // pick
		dfs(nums, i+1, n, cur, want) // don't pick
}

func countMaxOrSubsetsDP(nums []int) int {
	var dp [1 << 17]int
	var maxOR int
	dp[0] = 1
	for _, num := range nums {
		for a := maxOR; a >= 0; a-- {
			dp[a|num] += dp[a]
		}
		maxOR |= num
	}
	return dp[maxOR]
}

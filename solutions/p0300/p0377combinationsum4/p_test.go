package p0377combinationsum4

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum4(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, combinationSum4(tc.nums, tc.target))
		})
	}
}

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i-n < 0 || n > target {
				continue
			}
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

package p0198houserobber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rob(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		// {[]int{1, 2, 3, 1}, 4},
		// {[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, rob(tc.nums))
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := range nums {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}
	return dp[n+1]
}

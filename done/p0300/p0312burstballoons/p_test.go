package p0312burstbaloons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCoins(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxCoins(tc.nums))
		})
	}
}

func maxCoins(nums []int) int {
	newNums := make([]int, 0, len(nums)+2)
	newNums = append(newNums, 1)
	for _, num := range nums {
		if num > 0 {
			newNums = append(newNums, num)
		}
	}
	newNums = append(newNums, 1)
	nums = newNums
	n := len(nums)

	// dp[i][j] = max coins from popping balloons between i and j
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// For each round, pop balloons of width k
	for k := 1; k < n; k++ {
		for left := 0; left+k < n; left++ {
			right := left + k
			// Find the max score from [left,right]
			// Assume that a balloon at position i is the last balloon to burst
			// Since it is the last balloon, the max score from its left and right
			// side is sure to be independently solvable.
			for i := left + 1; i < right; i++ {
				dp[left][right] = max(dp[left][right],
					nums[left]*nums[i]*nums[right]+dp[left][i]+dp[i][right],
				)
			}
		}
	}
	// For all other cases, the max amount is the max of popping either side
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

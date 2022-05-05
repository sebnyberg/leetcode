package p1746maximumsubarraysumafteroneoperation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumAfterOperation(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, -1, -4, -3}, 17},
		{[]int{1, -1, 1, 1, -1, -1, 1}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumAfterOperation(tc.nums))
		})
	}
}

func maxSumAfterOperation(nums []int) int {
	// Intuition:
	// * As long as nums is non-empty, there will always be a positive max sum
	//
	// Idea: use Kadane's algorithm, but run two versions:
	// 1. Regular Kadane's without any squares
	// 2. Kadane's where we either square current or simply copy previous result
	//
	// Result is the max result in case 2
	sumNoSquare := nums[0]
	sumWithSquare := nums[0] * nums[0]
	maxSum := sumWithSquare
	n := len(nums)
	for i := 1; i < n; i++ {
		sumWithSquare = max(
			nums[i]*nums[i],
			max(
				sumNoSquare+nums[i]*nums[i],
				sumWithSquare+nums[i],
			),
		)
		if sumNoSquare > 0 {
			sumNoSquare += nums[i]
		} else {
			sumNoSquare = nums[i]
		}
		maxSum = max(maxSum, sumWithSquare)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p0918maximumsumcircularsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubarraySumCircular(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-2, -3, -1}, -1},
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{3, -1, 2, -1}, 4},
		{[]int{3, -2, 2, -3}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSubarraySumCircular(tc.nums))
		})
	}
}

func maxSubarraySumCircular(nums []int) int {
	// First case is simply the max sum of elements in the nums array
	maxSum := nums[0]
	curMax := 0
	var total int
	for _, n := range nums {
		curMax = max(curMax+n, n)
		maxSum = max(maxSum, curMax)
		total += n
	}

	// Second case finds the minimum subarray sum, which removed from nums would
	// form the maximum circular subarray sum.
	minSum := 0
	curMin := 0
	for _, n := range nums {
		curMin = min(curMin+n, n)
		minSum = min(minSum, curMin)
	}

	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	return maxSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

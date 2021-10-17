package p2036maxalternatingsubarraysum

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumAlternatingSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{-1}, -1},
		{[]int{3, -1, 1, 2}, 5},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumAlternatingSubarraySum(tc.nums))
		})
	}
}

func maximumAlternatingSubarraySum(nums []int) int64 {
	maxSum := nums[0]

	// Transform nums so that every odd element is negative
	// This transform is valid for any sum that starts in an even position.
	for i := 1; i < len(nums); i += 2 {
		nums[i] = -nums[i]
	}
	maxSum = max(maxSum, maxSumResetOnEven(nums))

	// Invert nums again so that every even element is negative.
	// This is valid for any sum that starts at an odd position
	for i := 0; i < len(nums); i++ {
		nums[i] = -nums[i]
	}
	maxSum = max(maxSum, maxSumResetOnEven(nums[1:])) // note! re-slice

	return int64(maxSum)
}

// maxSumResetOnEven returns the maximum sum in nums given that the sum can reset
// on even indices.
func maxSumResetOnEven(nums []int) int {
	maxSum := math.MinInt32
	var sum int
	for i, num := range nums {
		if i%2 == 0 {
			sum = max(sum, 0)
		}
		sum += num
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

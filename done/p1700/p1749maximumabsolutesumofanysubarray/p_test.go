package p1749maximumabsolutesumofanysubarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAbsoluteSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-7, -1, 0, -2, 1, 3, 8, -2, -6, -1, -10, -6, -6, 8, -4, -9, -4, 1, 4, -9}, 44},
		{[]int{1, -3, 2, 3, -4}, 5},
		{[]int{2, -5, 1, -4, 3, -2}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxAbsoluteSum(tc.nums))
		})
	}
}

func maxAbsoluteSum(nums []int) int {
	// Max absolute sum of any subarray is either -(minSum) or maxSum
	curMinSum := 0
	minSum := math.MaxInt32
	curMaxSum := 0
	maxSum := math.MinInt32
	for _, n := range nums {
		curMinSum = min(curMinSum+n, n)
		minSum = min(minSum, curMinSum)
		curMaxSum = max(curMaxSum+n, n)
		maxSum = max(maxSum, curMaxSum)
	}
	return max(-minSum, maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p0644maximumaveragesubarray2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxAverage(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.InEpsilon(t, tc.want, findMaxAverage(tc.nums, tc.k), 0.01)
		})
	}
}

const eps = 0.000_01

func findMaxAverage(nums []int, k int) float64 {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, n := range nums {
		minVal = min(minVal, n)
		maxVal = max(maxVal, n)
	}
	lower, upper := float64(minVal), float64(maxVal)
	for upper-lower > eps {
		mid := (upper + lower) / 2
		if isLargerThanMax(nums, k, mid) {
			lower = mid
		} else {
			upper = mid
		}
	}
	return upper
}

func isLargerThanMax(nums []int, k int, mid float64) bool {
	presum := 0.0
	cursum := 0.0
	for i := 0; i < k; i++ {
		cursum += float64(nums[i]) - mid
	}
	if cursum > 0 {
		return true
	}
	for i := k; i < len(nums); i++ {
		cursum += float64(nums[i]) - mid
		presum += float64(nums[i-k]) - mid
		if presum < 0 {
			cursum -= presum
			presum = 0.0
		}
		if cursum > 0 {
			return true
		}
	}
	return false
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

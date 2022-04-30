package p1

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumAverageDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 1, 1}, 0},
		{[]int{2, 5, 3, 9, 5, 3}, 3},
		{[]int{0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumAverageDifference(tc.nums))
		})
	}
}

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i := range nums {
		if i > 0 {
			presum[i] += presum[i-1]
		}
		presum[i] += nums[i]
	}
	minDiff := math.MaxInt64
	minDiffIdx := -1
	var rightSum int
	for i := n - 1; i >= 0; i-- {
		leftAvg := presum[i] / (i + 1)
		var rightAvg int
		if i < n-1 {
			r := n - 1 - i
			rightAvg = rightSum / r
		}
		if d := abs(rightAvg - leftAvg); d <= minDiff {
			minDiff = d
			minDiffIdx = i
		}
		rightSum += nums[i]
	}
	return minDiffIdx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

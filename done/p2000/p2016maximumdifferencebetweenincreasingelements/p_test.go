package p2016maximumdifferencebetweenincreasingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{7, 1, 5, 4}, 4},
		{[]int{9, 4, 3, 2}, -1},
		{[]int{1, 5, 2, 10}, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumDifference(tc.nums))
		})
	}
}

func maximumDifference(nums []int) int {
	maxDiff := -1
	minValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > minValue {
			maxDiff = max(maxDiff, nums[i]-minValue)
		}
		minValue = min(minValue, nums[i])
	}
	return maxDiff
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

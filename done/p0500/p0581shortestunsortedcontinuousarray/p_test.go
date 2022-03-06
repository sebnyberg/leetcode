package p0581shortestunsortedcontinuousarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findUnsortedSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 3, 3}, 2},
		{[]int{1, 3, 2, 2, 2}, 4},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findUnsortedSubarray(tc.nums))
		})
	}
}

func findUnsortedSubarray(nums []int) int {
	sorted := true
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	max := math.MinInt32
	var end int
	for i := range nums {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
	}
	min := math.MaxInt32
	var start int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			start = i
		}
	}
	return end - start + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

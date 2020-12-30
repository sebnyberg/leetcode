package l0016_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumClosest(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		// {[]int{-1, 2, 1, -4}, 1, 2},
		// {[]int{0, 1, 2}, 3, 3},
		{[]int{0, 1, 2}, 0, 3},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, threeSumClosest(tc.nums, tc.target))
		})
	}
}

// Find the triplet for which the sum minimizes the distance to the target
func threeSumClosest(nums []int, target int) int {
	// Sort numbers
	sort.Ints(nums)
	minsum := 9000

	for i, n := range nums {
		// Skip equal numbers
		if i > 0 && n == nums[i-1] {
			continue
		}
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := n + nums[lo] + nums[hi]
			switch {
			case sum == target:
				return target
			case sum > target:
				hi--
			case sum < target:
				lo++
			}
			if diff(sum, target) < diff(minsum, target) {
				minsum = sum
			}
		}
	}
	return minsum
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

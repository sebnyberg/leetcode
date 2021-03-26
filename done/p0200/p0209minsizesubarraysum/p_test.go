package p0209minsizesubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSubArrayLen(t *testing.T) {
	for _, tc := range []struct {
		target int
		nums   []int
		want   int
	}{
		{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 8},
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{11, []int{1, 2, 3, 4, 5}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, minSubArrayLen(tc.target, tc.nums))
		})
	}
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	var l, r int
	var sum int
	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		if sum >= target {
			break
		}
	}
	if sum < target {
		return 0
	}
	minLen := r - l + 1

	for {
		// Move left cursor until the sum is no longer contained
		for sum >= target {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}

		if r == n-1 {
			return minLen
		}

		// Move right until it can't move any more
		for r < n-1 && sum < target {
			r++
			sum += nums[r]
		}
		minLen = min(minLen, r-l+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

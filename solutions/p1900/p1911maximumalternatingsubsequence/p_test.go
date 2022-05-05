package p1911maximumalternatingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAlternatingSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{4, 2, 5, 3}, 7},
		{[]int{5, 6, 7, 8}, 8},
		{[]int{6, 2, 1, 2, 4, 5}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxAlternatingSum(tc.nums))
		})
	}
}

func maxAlternatingSum(nums []int) int64 {
	// Since there are up to 10k elements in total, O(n^2) is roughly OK but
	// exercise is likely to want O(nlogn) or O(n)
	// Cannot prove this but let's assume that the max alternating sum is
	// achieved by picking peaks and valleys in nums.

	// Lookahead
	i := 0
	n := len(nums)
	var res int
	for i < n-1 {
		// Find peak
		for i < n-1 && nums[i+1] >= nums[i] {
			i++
		}
		res += nums[i]
		if i == n {
			break
		}

		// Find valley
		for i < n-1 && nums[i+1] < nums[i] {
			i++
		}
		if i < n-1 {
			res -= nums[i]
		}
	}
	return int64(res)
}

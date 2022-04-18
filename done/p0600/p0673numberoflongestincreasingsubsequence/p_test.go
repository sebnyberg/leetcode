package p0673numberoflongestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findNumberOfLIS(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 4, 3, 5, 4, 7, 2}, 3},
		{[]int{1, 3, 5, 4, 7}, 2},
		{[]int{2, 2, 2, 2, 2}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findNumberOfLIS(tc.nums))
		})
	}
}

func findNumberOfLIS(nums []int) int {
	// Perform somewhat greedy LIS / patience sort. When growing beyond the
	// maximum length so far, reset count to 1. When changing the final value in
	// the stack, increment counter
	n := len(nums)
	maxlen := make([]int, n)
	maxcount := make([]int, n)
	for i := range maxlen {
		maxlen[i] = 1
		maxcount[i] = 1
	}
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			if maxlen[j]+1 > maxlen[i] {
				maxlen[i] = maxlen[j] + 1
				maxcount[i] = maxcount[j]
			} else if maxlen[j]+1 == maxlen[i] {
				maxcount[i] += maxcount[j]
			}
		}
	}

	var res int
	var maxLength int
	for _, v := range maxlen {
		maxLength = max(maxLength, v)
	}
	for i := range maxcount {
		if maxlen[i] == maxLength {
			res += maxcount[i]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

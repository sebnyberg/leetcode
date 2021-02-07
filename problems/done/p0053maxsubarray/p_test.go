package p0053maxsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		// {[]int{1, -1, 1}, 1},
		// {[]int{-1, -2}, -1},
		{[]int{-2, -1}, -1},
		// {[]int{-2, 1}, 1},
		// {[]int{-1, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		// {[]int{1}, 1},
		// {[]int{0}, 0},
		// {[]int{-1}, -1},
		// {[]int{-100000}, -100000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSubArray(tc.nums))
		})
	}
}

func maxSubArray(nums []int) int {
	n := len(nums)
	m := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

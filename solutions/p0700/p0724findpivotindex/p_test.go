package p0724findpivotindex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pivotIndex(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-1, -1, 0, 1, 1, 0}, 5},
		{[]int{-1, -1, 0, 0, -1, -1}, 2},
		{[]int{2, 1, -1}, 0},
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, pivotIndex(tc.nums))
		})
	}
}

func pivotIndex(nums []int) int {
	var rsum int
	for i := len(nums) - 1; i > 0; i-- {
		rsum += nums[i]
	}

	var lsum int
	for i, n := range nums[:len(nums)-1] {
		if lsum == rsum {
			return i
		}
		rsum -= nums[i+1]
		lsum += n
	}
	if lsum == 0 {
		return len(nums) - 1
	}
	return -1
}

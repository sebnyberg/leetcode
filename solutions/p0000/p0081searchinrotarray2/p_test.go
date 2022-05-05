package p0081searchinrotarray2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{3, 1}, 1, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{8, 8, 9, 9, 11, 12, 0, 2, 4, 6}, 3, false},
		{[]int{8, 8, 9, 9, 11, 12, 0, 2, 4, 6}, 8, true},
		{[]int{8, 8, 9, 9, 11, 12, 0, 2, 4, 6}, 7, false},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, search(tc.nums, tc.target))
		})
	}
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		// Pick mid point
		m := l + (r-l)/2
		if target == nums[m] {
			return true
		}
		switch {
		case nums[l] < nums[m]: // left side is sorted
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		case nums[m] < nums[r]: // right side is sorted
			if target <= nums[r] && target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		default: // not possible to know whether either side is sorted
			if target == nums[r] {
				return true
			}
			r--
		}
	}
	return false
}

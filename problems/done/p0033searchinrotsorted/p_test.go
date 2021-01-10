package p0033searchinrotsorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		// {[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		// {[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		// {[]int{1}, 0, -1},
		// {[]int{3, 1}, 1, 1},
		// {[]int{1, 3}, 2, -1},
		// {[]int{1, 3, 5}, 2, -1},
		{[]int{5, 1, 3}, 3, 2},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, search(tc.nums, tc.target))
		})
	}
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target < nums[l] || target > nums[mid] {
				l = mid + 1
				continue
			}
			r = mid - 1
			continue
		}
		// rhs is sorted
		if target < nums[mid] || target > nums[r] {
			r = mid - 1
			continue
		}
		l = mid + 1
	}

	return -1
}

package p0704binarysearch

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
		{[]int{5}, 5, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, search(tc.nums, tc.target))
		})
	}
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1

	// // Library version
	// idx := sort.SearchInts(nums, target)
	// if idx == len(nums) || nums[idx] != target {
	// 	return -1
	// }
	// return idx
}

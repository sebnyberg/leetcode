package p0034firstandlastinsorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchRange(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},
		{[]int{2, 2}, 3, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, searchRange(tc.nums, tc.target))
		})
	}
}

// Solution using std lib
// func searchRange(nums []int, target int) []int {
// 	res := make([]int, 2)
// 	n := len(nums)
// 	i := sort.SearchInts(nums, target)
// 	if i == n || nums[i] != target {
// 		return res
// 	}
// 	j := sort.SearchInts(nums[i:], target+1)
// 	res[0], res[1] = i, i+j-1
// 	return res
// }

func searchRange(nums []int, target int) []int {
	n := len(nums)
	res := make([]int, 2)
	res[0], res[1] = -1, -1
	if n == 0 {
		return res
	}

	l, r := 0, n
	for l < r {
		mid := int(l+r) >> 1
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		r = mid
	}
	if l == n || nums[l] != target {
		return res
	}

	for ; r < n && nums[r] == nums[l]; r++ {
	}

	res[0], res[1] = l, r-1
	return res
}

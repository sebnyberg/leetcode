package p0035searchinsert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchInsert(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, searchInsert(tc.nums, tc.target))
		})
	}
}

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) >> 1
		if nums[mid] < target {
			i = mid + 1
			continue
		}
		j = mid
	}
	return i
}

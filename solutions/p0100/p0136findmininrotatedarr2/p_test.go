package p0136findmininrotatedarr2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMin(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 0, 0, 0}, 0},
		{[]int{0, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{[]int{1, 0, 1, 1, 1, 1, 1, 1, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMin(tc.nums))
		})
	}
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (r + l) / 2
		if nums[m] > nums[r] {
			l = m + 1
			continue
		}
		if nums[m] == nums[r] && nums[m] == nums[l] {
			l++
			continue
		}

		// nums[m] <= nums[l]
		r = m
	}
	return nums[l]
}

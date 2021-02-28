package p0135findmininrotatedarr

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
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
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
		r = m
	}
	return nums[l]
}

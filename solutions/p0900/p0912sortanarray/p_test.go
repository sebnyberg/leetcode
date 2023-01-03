package p0912sortanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortArray(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 1, 3}, []int{1, 2, 3, 5}},
		{[]int{-1, 2, -8, 10}, []int{-8, -1, 2, 10}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sortArray(tc.nums))
		})
	}
}

func sortArray(nums []int) []int {
	n := len(nums)
	buf := make([]int, n)
	res := sort(nums, buf)
	return res
}

func sort(nums, buf []int) []int {
	// Let's do merge-sort
	if len(nums) == 1 {
		return nums
	}
	n := len(nums)
	left := sort(nums[:n/2], buf)
	right := sort(nums[n/2:], buf)
	var j int
	for ; len(left) > 0 && len(right) > 0; j++ {
		if left[0] < right[0] {
			buf[j] = left[0]
			left = left[1:]
		} else {
			buf[j] = right[0]
			right = right[1:]
		}
	}
	if len(left) > 0 {
		right = left
	}
	for ; len(right) > 0; j++ {
		buf[j] = right[0]
		right = right[1:]
	}
	copy(nums, buf[:len(nums)])
	return nums
}

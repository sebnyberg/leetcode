package p0795numberofsubarrayswithboundedmaximum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubarrayBoundedMax(t *testing.T) {
	for _, tc := range []struct {
		nums        []int
		left, right int
		want        int
	}{
		{[]int{2, 1, 4, 3}, 2, 3, 3},
		{[]int{2, 9, 2, 5, 6}, 2, 8, 7},
		{[]int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, 32, 69, 22},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numSubarrayBoundedMax(tc.nums, tc.left, tc.right))
		})
	}
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// whenever a number is > right, we cancel all counts
	// whenever a number is >= left, we increment the number of valid prior ranges
	// whenever a number is < left, we don't increment the number of valid prior
	// ranges, but we add the number of valid prior ranges to the answer
	prevValid, prevInvalid := -1, -1
	var res int
	for i := range nums {
		if nums[i] > right {
			prevValid = i
			prevInvalid = i
			continue
		}
		if nums[i] < left {
			res += prevValid - prevInvalid
			continue
		}
		prevValid = i
		res += i - prevInvalid
	}
	return res
}

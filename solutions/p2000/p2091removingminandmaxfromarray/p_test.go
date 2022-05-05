package p2091removingminandmaxfromarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDeletions(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-14, 61, 29, -18, 59, 13, -67, -16, 55, -57, 7, 74}, 6},
		{[]int{0, -4, 19, 1, 8, -2, -3, 5}, 3},
		{[]int{2, 10, 7, 5, 4, 1, 8, 6}, 5},
		{[]int{101}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDeletions(tc.nums))
		})
	}
}

func minimumDeletions(nums []int) int {
	var minIdx, maxIdx int
	minVal, maxVal := nums[0], nums[0]
	for i, num := range nums {
		if num < minVal {
			minVal = num
			minIdx = i
		}
		if num > maxVal {
			maxVal = num
			maxIdx = i
		}
	}
	first, second := minIdx, maxIdx
	if first > second {
		first, second = second, first
	}
	n := len(nums)
	res := second + 1
	res = min(res, (n-second)+first+1)
	res = min(res, n-first)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

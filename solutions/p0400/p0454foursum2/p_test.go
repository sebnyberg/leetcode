package p0454foursum2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fourSumCount(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
		want  int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, fourSumCount(tc.nums1, tc.nums2, tc.nums3, tc.nums4))
		})
	}
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// This count can be achieved by finding all possible sums for a pair of two
	// on both sides, then combining the 2sums to find the number of 4sum zeroes.
	n := len(nums1)
	left := make(map[int]int, n)
	right := make(map[int]int, n)
	for i := range nums1 {
		for j := range nums2 {
			left[nums1[i]+nums2[j]]++
			right[nums3[i]+nums4[j]]++
		}
	}
	var res int
	for leftSum, count := range left {
		res += right[-leftSum] * count
	}
	return res
}

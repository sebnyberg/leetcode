package p0496nextgreaterelementi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElement(t *testing.T) {
	for _, tc := range []struct {
		nums1, nums2 []int
		want         []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, nextGreaterElement(tc.nums1, tc.nums2))
		})
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numIdx := make(map[int]int)
	for i, num := range nums2 {
		numIdx[num] = i
	}
	n := len(nums2)
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		var j int
		for j = numIdx[num] + 1; j < n && nums2[j] <= num; j++ {
		}
		if j == n {
			res[i] = -1
		} else {
			res[i] = nums2[j]
		}
	}
	return res
}

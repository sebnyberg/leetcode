package p0496nextgreaterelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElement(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, nextGreaterElement(tc.nums1, tc.nums2))
		})
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2Idx := make(map[int]int)
	for i := range nums2 {
		nums2Idx[nums2[i]] = i
	}
	res := make([]int, len(nums1))
	for i := range nums1 {
		nums1[i] = -1
		if j, exists := nums2Idx[nums1[i]]; exists {
			for k := j + 1; k < len(nums2); k++ {
				if nums2[k] > nums1[i] {
					res[i] = nums2[k]
					break
				}
			}
		}
	}
	return res
}

package p1874minimizeproductsumoftwoarrays

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minProductSum(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{5, 3, 4, 2}, []int{4, 2, 2, 5}, 40},
		{[]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}, 65},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minProductSum(tc.nums1, tc.nums2))
		})
	}
}

func minProductSum(nums1 []int, nums2 []int) int {
	// Pretty sure this is just the sorted orders of nums1 and nums2...
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] > nums2[j] })
	var res int
	for i := range nums1 {
		res += nums1[i] * nums2[i]
	}
	return res
}

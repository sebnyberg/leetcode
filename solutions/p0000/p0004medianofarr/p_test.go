package p0004medianofarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMedianSortedArrays(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, findMedianSortedArrays(tc.nums1, tc.nums2))
		})
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	n := n1 + n2

	var l, r int
	next := func() int {
		if r == n2 || (l < n1 && nums1[l] <= nums2[r]) {
			res := nums1[l]
			l++
			return res
		}
		res := nums2[r]
		r++
		return res
	}

	// 4 elements => skip 1
	// 3 elements => skip 1
	// This gives the number of elements to skip: ((n+1)/2)-1
	for i := 0; i < ((n+1)/2)-1; i++ {
		next()
	}
	if n&1 == 1 {
		return float64(next())
	}
	return float64(next()+next()) / 2
}

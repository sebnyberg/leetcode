package p0349intersectionoftwoarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersection(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, intersection(tc.nums1, tc.nums2))
		})
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	var seen [1001]bool
	for _, num := range nums1 {
		seen[num] = true
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		if seen[num] {
			res = append(res, num)
			seen[num] = false
		}
	}
	return res
}

package p0350intersectionoftwoarrays2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersect(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, intersect(tc.nums1, tc.nums2))
		})
	}
}

func intersect(nums1 []int, nums2 []int) []int {
	var count [1001]int
	for _, n := range nums1 {
		count[n]++
	}
	res := make([]int, 0)
	for _, n := range nums2 {
		if count[n] > 0 {
			count[n]--
			res = append(res, n)
		}
	}
	return res
}

package p2032twooutofthree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoOutOfThree(t *testing.T) {
	for _, tc := range []struct {
		nums1, nums2, nums3 []int
		want                []int
	}{
		{[]int{1, 1, 3, 2}, []int{2, 3}, []int{3}, []int{3, 2}},
		{[]int{3, 1}, []int{2, 3}, []int{1, 2}, []int{2, 3, 1}},
		{[]int{1, 2, 2}, []int{4, 3, 3}, []int{5}, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, twoOutOfThree(tc.nums1, tc.nums2, tc.nums3))
		})
	}
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var seen [101][3]bool
	for i := range nums1 {
		seen[nums1[i]][0] = true
	}
	for i := range nums2 {
		seen[nums2[i]][1] = true
	}
	for i := range nums3 {
		seen[nums3[i]][2] = true
	}
	res := make([]int, 0)
	for num := range seen {
		var count int
		for arrIdx := range seen[num] {
			if seen[num][arrIdx] {
				count++
			}
		}
		if count >= 2 {
			res = append(res, num)
		}
	}
	return res
}

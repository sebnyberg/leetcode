package p0088mergesortedarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{0, 0, 0}, 0, []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{3, 4, 5, 0, 0}, 3, []int{1, 6}, 2, []int{1, 3, 4, 5, 6}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v/%v/%+v/%v", tc.nums1, tc.m, tc.nums2, tc.n), func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			require.Equal(t, tc.want, tc.nums1)
		})
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; i < m+n; {
		if i == m {
			copy(nums1[i+j:], nums2[j:])
			return
		}
		if j == n {
			return
		}
		if i < m && nums1[i+j] <= nums2[j] {
			i++
			continue
		}
		if j < n && nums2[j] < nums1[i+j] {
			copy(nums1[i+j+1:], nums1[i+j:]) // make space
			nums1[i+j] = nums2[j]
			j++
		}
	}
}

package p1855maximumdistancebetweenpairofvalues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDistance(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}, 2},
		{[]int{2, 2, 2}, []int{10, 10, 1}, 1},
		{[]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}, 2},
		{[]int{5, 4}, []int{3, 2}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, maxDistance(tc.nums1, tc.nums2))
		})
	}
}

func maxDistance(nums1 []int, nums2 []int) int {
	var maxdist int
	n2 := len(nums2)
	for i, n := range nums1 {
		if n2-i <= maxdist {
			break
		}
		for j := n2 - 1; j > i; j-- {
			if nums2[j] >= n {
				maxdist = max(maxdist, j-i)
				break
			}
		}
	}
	return maxdist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

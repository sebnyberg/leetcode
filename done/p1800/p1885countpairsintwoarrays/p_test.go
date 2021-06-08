package p1885countpairsintwoarrays

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPairs(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int64
	}{
		{[]int{1, 10, 6, 2}, []int{1, 4, 1, 5}, 5},
		{[]int{2, 1, 2, 1}, []int{1, 2, 1, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, countPairs(tc.nums1, tc.nums2))
		})
	}
}

// countPairs counts the pairs of indices (i, j) in nums1, nums2 such that
// i < j and nums1[i] + nums1[j] > nums2[i] + nums2[j]
func countPairs(nums1 []int, nums2 []int) int64 {
	// nums1[i] + nums1[j] > nums2[i] + nums2[j]
	// <=>
	// nums1[i] - nums2[i] > nums2[j] - nums1[j]
	// ==(nums1[i] -= nums2[i])=>
	// nums1[i] > -nums1[j]
	// <=>
	// nums1[i] + nums1[j] > 0 (is true for all j where nums1[j] > -nums1[i])
	n := len(nums1)
	for i, num := range nums2 {
		nums1[i] -= num
	}

	var count int
	sort.Ints(nums1)
	for i := range nums1 {
		pos := sort.SearchInts(nums1[i+1:], -nums1[i]+1) + i + 1
		if pos != n {
			count += n - pos
		}
	}
	return int64(count)
}

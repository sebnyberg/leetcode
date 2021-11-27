package p1983widestpairofindiceswithequalrangesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_widestPairOfIndices(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{0, 1, 1}, []int{1, 0, 1}, 3},
		{[]int{0, 1}, []int{1, 1}, 1},
		{[]int{1, 1, 0, 1}, []int{0, 1, 1, 0}, 3},
		{[]int{0}, []int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, widestPairOfIndices(tc.nums1, tc.nums2))
		})
	}
}

func widestPairOfIndices(nums1 []int, nums2 []int) int {
	// This is a typical transformation exercise.
	// If nums2 is subtracted from nums1, then the exercise is instead to find
	// the widest pair of indices in nums such that its sum is zero
	//
	// For any prefix sum, it is "optimal" if its value is unique and its index
	// is lower than any other indices in nums1.
	for i := range nums2 {
		nums1[i] -= nums2[i]
	}
	sumIdx := make(map[int]int)
	sumIdx[0] = 0
	var sum int
	var maxWidth int
	for i := range nums1 {
		sum += nums1[i]
		if idx, exists := sumIdx[sum]; exists {
			if dist := i + 1 - idx; dist > maxWidth {
				maxWidth = dist
			}
		} else {
			sumIdx[sum] = i + 1
		}
	}
	return maxWidth
}

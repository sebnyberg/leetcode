package p0004medianofarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_medianOfArr(t *testing.T) {
	tcs := []struct {
		in   []int
		want float64
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 0.5},
		{[]int{0, 1, 2}, 1},
		{[]int{0, 1, 2, 3}, 1.5},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, medianOfArr(tc.in))
		})
	}
}

func medianOfArr(a []int) float64 {
	var alen int = len(a)
	switch alen {
	case 0:
		panic("median not allowed for zero-length array")
	case 1:
		return float64(a[0])
	}
	if alen%2 == 0 {
		upperMid := alen / 2
		return float64(a[upperMid-1]+a[upperMid]) / 2
	}
	return float64(a[alen/2])
}

func Test_findMedianSortedArrays(t *testing.T) {
	tcs := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.nums1, tc.nums2), func(t *testing.T) {
			require.Equal(t, tc.want, findMedianSortedArrays(tc.nums1, tc.nums2))
		})
	}
}

func Test_mergeArraysSorted(t *testing.T) {
	tcs := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{0}, []int{0}, []int{0, 0}},
		{nil, []int{0}, []int{0}},
		{[]int{0}, nil, []int{0}},
		{[]int{0, 1}, []int{0}, []int{0, 0, 1}},
		{[]int{0, 1, 2}, []int{0}, []int{0, 0, 1, 2}},
		{[]int{0, 1, 2}, []int{0, 1, 2}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{0}, []int{0, 1, 2}, []int{0, 0, 1, 2}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.nums1, tc.nums2), func(t *testing.T) {
			require.Equal(t, tc.want, mergeArraysSorted(tc.nums1, tc.nums2))
		})
	}
}

func mergeArraysSorted(nums1 []int, nums2 []int) []int {
	// Edge-cases
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}

	// Merged, sorted array
	merged := make([]int, len(nums1)+len(nums2))

	// i => nums1 index
	// j => nums2 index
	var i, j int

	for {
		// Exit condition
		if i == len(nums1) && j == len(nums2) {
			break
		}

		// Edge cases
		// nums1 done, iterate j
		if i == len(nums1) {
			merged[i+j] = nums2[j]
			j++
			continue
		}
		// nums2 is done or nums1 has a smaller/equal number, iterate i
		if j == len(nums2) || nums1[i] <= nums2[j] {
			merged[i+j] = nums1[i]
			i++
			continue
		}

		merged[i+j] = nums2[j]
		j++
	}

	return merged
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge arrays into sorted list
	merged := mergeArraysSorted(nums1, nums2)

	return medianOfArr(merged)
}

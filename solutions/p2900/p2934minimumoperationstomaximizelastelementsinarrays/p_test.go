package p2934minimumoperationstomaximizelastelementsinarrays

func minOperations(nums1 []int, nums2 []int) int {
	// Let's see here.. we can verify that an answer exists by checking that
	// there exists a number for each position in nums1 and nums2 such that the
	// numbers are <= the last number.
	n := len(nums1)
	a := nums1[n-1]
	b := nums2[n-1]
	if b < a {
		b, a = a, b
	}
	for i := range nums1 {
		x, y := nums1[i], nums2[i]
		if y < x {
			x, y = y, x
		}
		if x > a || y > b {
			return -1
		}
	}
	// A solution exists, the question is how to minimize the number of swaps.
	// swaps1 == swaps if keeping the original max values at the end of nums
	// swaps2 == swaps if changing max values at the end of nums
	var swaps1, swaps2 int

	for i := range nums1 {
		if nums1[i] > nums1[n-1] || nums2[i] > nums2[n-1] {
			swaps1++
		}
		if nums1[i] > nums2[n-1] || nums2[i] > nums1[n-1] {
			swaps2++
		}
	}
	return min(swaps1, swaps2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

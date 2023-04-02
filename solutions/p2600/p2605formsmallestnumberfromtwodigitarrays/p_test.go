package p2605formsmallestnumberfromtwodigitarrays

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	// Try to find one match
	var a [10]bool
	for _, x := range nums1 {
		a[x] = true
	}
	var b [10]bool
	for _, x := range nums2 {
		b[x] = true
	}
	for i := range a {
		if a[i] && b[i] {
			return i
		}
	}
	nums := []int{}
	for i := range a {
		if a[i] {
			nums = append(nums, i)
			break
		}
	}
	for i := range b {
		if b[i] {
			nums = append(nums, i)
			break
		}
	}
	sort.Ints(nums)
	return nums[0]*10 + nums[1]
}

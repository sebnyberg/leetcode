package p2294partitionarraysuchthatmaximumdifferenceisk

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	minVal := nums[0]
	count := 1
	for i := 1; i < n; i++ {
		if nums[i]-minVal <= k {
			continue
		}
		count++
		minVal = nums[i]
	}
	return count
}

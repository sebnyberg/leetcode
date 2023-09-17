package p2860happystudents

import "sort"

func countWays(nums []int) int {
	// Interestingly, there is either no group of a certain length or only
	// exactly one. The reason is that any given number is either smaller,
	// equal, or larger than the current length. This makes it either valid for
	// inclusion, makes it so that there are no solutions at all, or is valid
	// for exclusion. Two cases can never be true at the same time.
	sort.Ints(nums)
	var res int
	for i := range nums {
		if nums[i] > i && (i == 0 || nums[i-1] < i) {
			res++
		}
	}
	if nums[len(nums)-1] < len(nums) {
		res++
	}
	return res
}

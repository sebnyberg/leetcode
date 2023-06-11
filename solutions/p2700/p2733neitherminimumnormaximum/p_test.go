package p2733neitherminimumnormaximum

import "sort"

func findNonMinOrMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if len(nums) > 2 && nums[n-1]-nums[0] >= 2 {
		return nums[1]
	}
	return -1
}

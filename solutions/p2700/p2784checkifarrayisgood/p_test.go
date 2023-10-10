package p2784checkifarrayisgood

import "sort"

func isGood(nums []int) bool {
	n := len(nums) - 1
	sort.Ints(nums)
	for x := 1; x <= n; x++ {
		if nums[x-1] != x {
			return false
		}
	}
	return nums[n] == n
}

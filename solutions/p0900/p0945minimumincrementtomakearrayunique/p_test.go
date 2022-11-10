package p0945minimumincrementtomakearrayunique

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := range nums {
		if i > 0 && nums[i] <= nums[i-1] {
			res += (nums[i-1] - nums[i]) + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}

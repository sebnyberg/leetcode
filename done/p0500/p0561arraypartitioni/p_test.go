package p0561arraypartitioni

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

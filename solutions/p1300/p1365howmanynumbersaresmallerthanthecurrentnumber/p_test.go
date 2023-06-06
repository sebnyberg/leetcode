package p1365howmanynumbersaresmallerthanthecurrentnumber

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})
	res := make([]int, n)
	for i := range idx {
		if i > 0 && nums[idx[i-1]] == nums[idx[i]] {
			res[idx[i]] = res[idx[i-1]]
		} else {
			res[idx[i]] = i
		}
	}
	return res
}

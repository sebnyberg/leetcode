package p0628maximumproductofthreenumbers

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[n-3] * nums[n-2] * nums[n-1]
	res = max(res, nums[0]*nums[n-2]*nums[n-1])
	res = max(res, nums[0]*nums[1]*nums[n-1])
	res = max(res, nums[0]*nums[1]*nums[2])
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

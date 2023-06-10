package p1403minimumsubsequenceinnonincreasingorder

import "sort"

func minSubsequence(nums []int) []int {
	var sum int
	for _, x := range nums {
		sum += x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var res []int
	var resSum int
	for i := 0; resSum <= sum; i++ {
		resSum += nums[i]
		sum -= nums[i]
		res = append(res, nums[i])
	}
	return res
}

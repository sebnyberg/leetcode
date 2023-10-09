package p2779maximumbeautyofanarrayafterapplyingoperation

import "sort"

func maximumBeauty(nums []int, k int) int {
	// Omg I thought the question was about subARRAY
	//
	// Subsequence is easy as hell..
	//
	sort.Ints(nums)
	var j int
	var res int
	for i := range nums {
		for nums[i]-k > nums[j]+k {
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

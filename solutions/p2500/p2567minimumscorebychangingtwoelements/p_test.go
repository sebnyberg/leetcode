package p2567minimumscorebychangingtwoelements

import "sort"

func minimizeSum(nums []int) int {
	// The low score is the delta between the two closest elements.
	// The max score is the delta between the two most different elements
	//
	// How do we reduce this sum?
	//
	// 1. Change any element to be equal to another element. This reduces the
	// low score to zero. In fact, we can always redure the low score to zero as
	// we reduce the max score.
	//
	// 2. Change the largest/smallest number.
	//
	// This tells us that we can either:
	//
	// 1) Disregard largest two numbers,
	// 2) Disregard largest and smallest number,
	// 3) Disregard smallest two numbers
	//
	sort.Ints(nums)
	n := len(nums)
	res := nums[n-3] - nums[0]
	res = min(res, nums[n-2]-nums[1])
	res = min(res, nums[n-1]-nums[2])
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

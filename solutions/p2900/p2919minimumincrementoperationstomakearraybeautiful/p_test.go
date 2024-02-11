package p2919minimumincrementoperationstomakearraybeautiful

import "math"

func minIncrementOperations(nums []int, k int) int64 {
	// Clearly, the "nasty" case is a length-3 window that scans across nums,
	// trying to find a subarray that does not contain a maximum value of k
	//
	// Whenever we do encounter such a window, we have three options:
	//
	// 1. Increment first value until we meet the criteria,
	// 2. Increment second value until we meet the criteria,
	// 3. Increment the third value until we meet the criteria
	//
	var dp [3]int

	dp[0] = max(0, k-nums[0])
	dp[1] = max(0, k-nums[1])
	dp[2] = max(0, k-nums[2])

	for i := 3; i < len(nums); i++ {
		dp[0], dp[1], dp[2] = dp[1], dp[2], min(dp[0], min(dp[1], dp[2]))+max(0, k-nums[i])
	}
	return int64(min(dp[0], min(dp[1], dp[2])))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

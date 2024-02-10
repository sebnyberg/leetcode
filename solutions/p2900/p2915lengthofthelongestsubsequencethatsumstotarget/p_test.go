package p2915lengthofthelongestsubsequencethatsumstotarget

import "math"

func lengthOfLongestSubsequence(nums []int, target int) int {
	// Create a list of possible current sums where we try to have as many
	// elements as possible.
	var dp [1001]int
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = 0
	for i := range nums {
		for x := target; x-nums[i] >= 0; x-- {
			dp[x] = max(dp[x], dp[x-nums[i]]+1)
		}
	}
	if dp[target] < 0 {
		return -1
	}
	return dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1458maxdotproductoftwosubsequences

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {
	// This is a typical matching problem. We can use DP for it.
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			dp[i][j] = max(
				nums1[i-1]*nums2[j-1]+dp[i-1][j-1],
				max(dp[i][j-1], dp[i-1][j]),
			)
		}
	}
	a := math.MinInt32
	for _, x := range nums1 {
		for _, y := range nums2 {
			a = max(a, x*y)
		}
	}
	if a < 0 {
		return a
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

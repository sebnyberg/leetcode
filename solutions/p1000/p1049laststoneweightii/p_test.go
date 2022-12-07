package p4

import "math"

func lastStoneWeightII(stones []int) int {
	// This is a really tricky problem.
	//
	// The solution is the dual-knapsack problem, where we want to minimize the
	// sum difference between the two knapsacks.
	//
	// Why?
	//
	// Imagine that we split the stones into two groups.
	//
	// If all elements are equal, then that's our solution: 0.
	//
	// If one element is larger in one of the groups, then what will remain is
	// the delta between that element and some element in the other group. This
	// would be the optimal result of matching stones from the two groups.
	//
	// Now, you could easily construct a grouping that results in more than one
	// element remaining on each side at the end, but that would not be an
	// optimally split partition. In that scenario, one of the elements could've
	// been moved over to the other side.
	//
	var sum int
	for _, s := range stones {
		sum += s
	}
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, s := range stones {
		for i := sum; i >= s; i-- {
			dp[i] = dp[i] || dp[i-s]
		}
	}
	// Find optimal sum
	delta := math.MaxInt32
	for k := sum; k >= 0; k-- {
		if dp[k] {
			delta = min(delta, abs(k-(sum-k)))
		}
	}
	return delta
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package p2560houserobberiv

import "math"

func minCapability(nums []int, k int) int {
	// We wish to minimize the amount of money stolen such that the robber's
	// conditions are still satisfied.
	//
	// This is a classical DP problem, but it scales with O(n*k) which is too
	// large. So we need a different approach.
	//
	// I guess we could remove the most valuable houses one by one until the
	// remaining configuration is no longer valid? The problem then lies in
	// finding a data structure that can attach each index to its segment so
	// that the segments can continuously be split up.
	//
	// Another approach is to greedily visit every valid house, and validate
	// during binary search.
	//
	check := func(x int) bool {
		var m int
		var i int
		for i < len(nums) && m < k {
			if nums[i] <= x {
				m++
				i += 2
			} else {
				i++
			}
		}
		return m == k
	}

	lo, hi := 0, math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

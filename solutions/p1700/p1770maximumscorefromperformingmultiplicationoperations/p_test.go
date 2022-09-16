package p1770maximumscorefromperformingmultiplicationoperations

import "math"

func maximumScore(nums []int, multipliers []int) int {
	// Let's try a naive memoized (top-down) approach and see if it's ok.
	// The current state is the current left/right position in nums
	// Once the count covers multipliers, then we're done
	m := len(multipliers)
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, m)
		for j := range mem[i] {
			mem[i][j] = math.MinInt32
		}
	}

	res := dp(mem, 0, len(nums)-1, len(nums), nums, multipliers)
	return res
}

func dp(mem [][]int, l, r, n int, nums, multipliers []int) int {
	if len(multipliers) == 0 {
		return 0
	}
	v := mem[l][n-1-r]
	if v != math.MinInt32 {
		return v
	}
	// Try both options
	res := nums[l]*multipliers[0] + dp(mem, l+1, r, n, nums, multipliers[1:])
	res = max(res, nums[r]*multipliers[0]+dp(mem, l, r-1, n, nums, multipliers[1:]))
	mem[l][n-1-r] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p1402reducingdishes

import "sort"

func maxSatisfaction(satisfaction []int) int {
	n := len(satisfaction)
	sort.Ints(satisfaction)
	if satisfaction[n-1] <= 0 {
		return 0
	}
	curr := make([]int, n)
	next := make([]int, n)
	copy(curr, satisfaction)
	var res int
	for t := 2; t <= n; t++ {
		for i := t - 1; i < n; i++ {
			next[i] = curr[i-1] + satisfaction[i]*t
			res = max(res, next[i])
		}
		curr, next = next, curr
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

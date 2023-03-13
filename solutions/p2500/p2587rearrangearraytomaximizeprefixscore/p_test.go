package p2587rearrangearraytomaximizeprefixscore

import "sort"

func maxScore(nums []int) int {
	// Sort and greedily add as large numbers as possible until prefix becomes
	// non-positive.
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var sum int
	var res int
	for i := range nums {
		sum += nums[i]
		if sum <= 0 {
			break
		}
		res++
	}
	return res
}

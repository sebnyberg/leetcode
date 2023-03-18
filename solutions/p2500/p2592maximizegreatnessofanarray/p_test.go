package p2592maximizegreatnessofanarray

import "sort"

func maximizeGreatness(nums []int) int {
	// Greedy solution. The best matching is the smallest possible value larger
	// than its counterpart in nums. Sort and match using a two-pointer
	// approach.
	var l int
	var r int
	n := len(nums)
	sort.Ints(nums)
	var res int
	for l < n && r < n {
		for r < n && nums[r] <= nums[l] {
			r++
		}
		if r == n {
			break
		}
		res++
		l++
		r++
	}
	return res
}

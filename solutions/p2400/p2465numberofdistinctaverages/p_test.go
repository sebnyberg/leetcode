package p2465numberofdistinctaverages

import "sort"

func distinctAverages(nums []int) int {
	m := make(map[int]struct{})
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		sum := nums[i] + nums[len(nums)-1-i]
		m[sum] = struct{}{}
	}
	return len(m)
}

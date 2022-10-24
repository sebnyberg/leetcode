package p2448minimumcosttomakearrayequal

import "sort"

type numSorter struct {
	nums []int
	cost []int
}

func (s *numSorter) Swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	s.cost[i], s.cost[j] = s.cost[j], s.cost[i]
}

func (s numSorter) Len() int {
	return len(s.nums)
}

func (s numSorter) Less(i, j int) bool {
	return s.nums[i] < s.nums[j]
}

func minCost(nums []int, cost []int) int64 {
	s := &numSorter{
		nums: nums,
		cost: cost,
	}
	sort.Sort(s)
	n := len(nums)
	left := make([]int64, n+1)
	var acc int
	for i := 1; i < n; i++ {
		acc += cost[i-1]
		left[i] = left[i-1] + int64((nums[i]-nums[i-1])*acc)
	}
	var res int64
	res = left[n-1]
	acc = 0
	var right int64
	for i := n - 2; i >= 0; i-- {
		acc += cost[i+1]
		right += int64((nums[i+1] - nums[i]) * acc)
		res = min(res, right+left[i])
	}

	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

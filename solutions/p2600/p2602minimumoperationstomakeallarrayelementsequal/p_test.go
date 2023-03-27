package p2602minimumoperationstomakeallarrayelementsequal

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	n := len(nums)
	presum := make([]int, n+1)
	var sum int
	for i := range nums {
		presum[i+1] = nums[i] + presum[i]
		sum += nums[i]
	}

	m := len(queries)
	res := make([]int64, m)
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		// anything prior to nums[j] is <= q
		// anything at or after nums[j] is >= q
		suml := presum[j]
		sumr := sum - suml
		l := q*j - suml
		r := sumr - q*(n-j)
		res[i] = int64(l + r)
	}
	return res
}

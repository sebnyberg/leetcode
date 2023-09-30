package p2750waystosplitarrayintogoodsubarrays

import "math"

const mod = 1e9 + 7

func numberOfGoodSubarraySplits(nums []int) int {
	// I believe it is the product of the lengths of zeroes + 1 between ones.
	count := math.MinInt32
	var res int
	for i := range nums {
		if nums[i] == 0 {
			count++
			continue
		}
		if res == 0 {
			res = 1
		}
		if count > 0 {
			res = (res * (count + 1)) % mod
		}
		count = 0
	}
	return res
}

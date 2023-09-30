package p2740findthevalueofthepartition

import (
	"math"
	"sort"
)

func findValueOfPartition(nums []int) int {
	// The answer is just the smallest distance between two consecutive numbers
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

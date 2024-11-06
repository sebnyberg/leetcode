package p3011findifarraycanbesorted

import (
	"math"
	"math/bits"
)

func canSortArray(nums []int) bool {
	// Simply capture the max value of numbers with similar bitcounts
	x := bits.OnesCount(uint(nums[0]))
	maxVal := math.MinInt32
	nextMax := nums[0]
	for i := range nums {
		y := bits.OnesCount(uint(nums[i]))
		if y != x {
			maxVal = nextMax
			nextMax = nums[i]
			x = y
		}
		if nums[i] < maxVal {
			return false
		}
		nextMax = max(nextMax, nums[i])
	}
	return true
}

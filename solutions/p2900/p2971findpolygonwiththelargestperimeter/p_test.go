package p2971findpolygonwiththelargestperimeter

import (
	"fmt"
)

func largestPerimeter(nums []int) int64 {
	var sum int
	sort.Ints(nums)
	for _, x := range nums {
		sum += x
	}
	i := len(nums) - 1
	for ; i >= 2 && nums[i] >= sum-nums[i]; i-- {
		sum -= nums[i]
	}
	if i == 1 {
		return -1
	}
	return int64(sum)
}

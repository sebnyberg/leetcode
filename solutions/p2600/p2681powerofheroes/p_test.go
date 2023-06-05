package p2681powerofheroes

import "sort"

const mod = 1e9 + 7

func sumOfPower(nums []int) int {
	// Keep track of the sum of min values for all possible prior groups.
	// The contribution of the current (max) value to the total will be max^2 *
	// sum of min values for all possible prior groups. The new sum of min
	// values will be sum * 2 + nums[i]
	sort.Ints(nums)
	var res int
	var minSum int
	for i := range nums {
		sq := (nums[i] * nums[i]) % mod
		res = (res + (minSum * sq)) % mod
		res = (res + nums[i]*sq) % mod
		minSum = (minSum*2 + nums[i]) % mod
	}
	return res
}

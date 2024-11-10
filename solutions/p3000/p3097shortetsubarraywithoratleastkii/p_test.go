package p3097shortetsubarraywithoratleastkii

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	var count [32]int
	var l int
	var curr int
	res := math.MaxInt32
	for r, val := range nums {
		x := val
		for i := 0; x > 0; i++ {
			count[i] += x & 1
			if count[i] == 1 {
				curr |= 1 << i
			}
			x >>= 1
		}
		for curr >= k && l <= r {
			res = min(res, r-l+1)
			// remove nums[l]
			x := nums[l]
			for i := 0; x > 0; i++ {
				if count[i] == 1 && x&1 == 1 {
					curr &^= 1 << i
				}
				count[i] -= x & 1
				x >>= 1
			}
			l++
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

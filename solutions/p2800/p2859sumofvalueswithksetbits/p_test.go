package p2859sumofvalueswithksetbits

import "math/bits"

func sumIndicesWithKSetBits(nums []int, k int) int {
	var res int
	for i, x := range nums {
		if bits.OnesCount(uint(i)) == k {
			res += x
		}
	}
	return res
}

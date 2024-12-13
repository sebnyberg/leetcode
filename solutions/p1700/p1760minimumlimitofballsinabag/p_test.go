package p1716minimumlimitofballsinabag

import "math"

func minimumSize(nums []int, maxOperations int) int {
	// binary search
	ok := func(y int) bool {
		var res int
		for _, x := range nums {
			res += (x - 1) / y
			if res > maxOperations {
				return false
			}
		}
		return true
	}
	lo := 1
	hi := math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

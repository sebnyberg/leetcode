package p1283findthesmallestdivisorgivenathreshold

import "math"

func smallestDivisor(nums []int, threshold int) int {
	// Because the range of possible values is monotonically decreasing as the
	// divisor grows, we can use binary search.
	check := func(d int) bool {
		var res int
		for _, x := range nums {
			res += (x + (d - 1)) / d
		}
		return res <= threshold
	}

	lo, hi := 1, math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

package p1011capacitytoshippackageswithinddays

import "math"

func shipWithinDays(weights []int, days int) int {
	ok := func(w int) bool {
		day := 1
		var ww int
		for i := range weights {
			if weights[i] > w {
				return false
			}
			ww += weights[i]
			if ww > w {
				day++
				ww = weights[i]
			}
		}
		return day <= days
	}

	var lo int
	hi := math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

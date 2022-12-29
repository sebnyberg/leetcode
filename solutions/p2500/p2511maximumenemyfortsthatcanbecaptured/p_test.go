package p2511maximumenemyfortsthatcanbecaptured

import "math"

func captureForts(forts []int) int {
	var res int
	prev := math.MaxInt32
	for i, x := range forts {
		if x == 1 {
			prev = i
			continue
		}
		if x == -1 {
			res = max(res, i-prev-1)
			prev = math.MaxInt32
		}
	}

	prev = math.MinInt32
	for i := len(forts) - 1; i >= 0; i-- {
		x := forts[i]
		if x == 1 {
			prev = i
			continue
		}
		if x == -1 {
			res = max(res, prev-i-1)
			prev = math.MinInt32
		}
	}
	if res < 0 {
		return 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

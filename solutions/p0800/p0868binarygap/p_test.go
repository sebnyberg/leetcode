package p0868binarygap

import "math"

func binaryGap(n int) int {
	prev := math.MaxInt32
	var res int
	for i := 0; n > 0; i++ {
		if n&1 > 0 {
			res = max(res, i-prev)
			prev = i
		}
		n >>= 1
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p2453destroysequentialtargets

func destroyTargets(nums []int, space int) int {
	m := make(map[int]int)
	minVal := make(map[int]int)
	var res int
	var maxCount int
	for _, x := range nums {
		r := x % space
		m[r]++
		if v, exists := minVal[r]; !exists || v > x {
			minVal[r] = x
		}
		if m[r] > maxCount || m[r] == maxCount && minVal[r] < res {
			res = minVal[r]
		}
		maxCount = max(maxCount, m[r])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

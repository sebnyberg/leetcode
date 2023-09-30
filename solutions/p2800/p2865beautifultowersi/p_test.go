package p2865beautifultowersi

func maximumSumOfHeights(maxHeights []int) int64 {
	// Forgot to check constraints.. we can simply calculate the result for each
	// index.
	calculate := func(i int) int {
		max := maxHeights[i]
		res := max
		for j := i - 1; j >= 0; j-- {
			max = min(max, maxHeights[j])
			res += max
		}
		max = maxHeights[i]
		for j := i + 1; j < len(maxHeights); j++ {
			max = min(max, maxHeights[j])
			res += max
		}
		return res
	}
	var res int
	for i := range maxHeights {
		res = max(res, calculate(i))
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

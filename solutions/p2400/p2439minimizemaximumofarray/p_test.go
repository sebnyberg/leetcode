package p2439minimizemaximumofarray

func minimizeArrayValue(nums []int) int {
	var mean int
	var sum int
	for i, x := range nums {
		sum += x
		maybeNewMean := sum / (i + 1)
		if sum%(i+1) != 0 {
			maybeNewMean++
		}
		mean = max(mean, maybeNewMean)
	}
	return mean
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

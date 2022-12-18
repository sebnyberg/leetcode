package p2501longestsquarestreakinanarray

func longestSquareStreak(nums []int) int {
	seen := make(map[int]bool)
	for _, x := range nums {
		seen[x] = true
	}
	var maxCount int
	for x := 2; x <= 1e5; x++ {
		var count int
		for y := x; seen[y]; y = y * y {
			count++
		}
		if count > 1 {
			maxCount = max(maxCount, count)
		}
	}
	if maxCount == 0 {
		return -1
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

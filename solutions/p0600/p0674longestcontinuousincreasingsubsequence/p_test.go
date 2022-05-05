package p0674longestcontinuousincreasingsubsequence

func findLengthOfLCIS(nums []int) int {
	count := 1
	maxCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

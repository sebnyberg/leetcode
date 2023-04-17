package p2640findthescoreofallprefixesofanarray

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
	res := make([]int64, n)
	res[0] = int64(nums[0] + nums[0])
	maxVal := nums[0]
	var sum int
	for i := range nums {
		maxVal = max(maxVal, nums[i])
		sum += nums[i] + maxVal
		res[i] = int64(sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

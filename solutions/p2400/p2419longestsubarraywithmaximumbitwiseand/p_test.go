package p2419longestsubarrayrithbitwiseand

func longestSubarray(nums []int) int {
	// The answer is simply the longest subarray containing only the maximum
	// digit
	var maxVal, count int
	var res int
	for _, x := range nums {
		if x > maxVal {
			maxVal = x
			count = 1
			res = 1
		} else if x < maxVal {
			count = 0
		} else {
			count++
		}
		res = max(res, count)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

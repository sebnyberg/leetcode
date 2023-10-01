package p2874maximumvalueofanorderedtripletii

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	maxRight := make([]int, n+1)
	for i := len(nums) - 1; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], nums[i])
	}
	maxLeft := nums[0]
	var res int
	for i := 1; i < len(nums)-1; i++ {
		res = max(res, (maxLeft-nums[i])*maxRight[i+1])
		maxLeft = max(maxLeft, nums[i])
	}
	return int64(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

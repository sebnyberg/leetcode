package p2789largestelementinanarrayaftermergeoperations

func maxArrayValue(nums []int) int64 {
	var res int
	x := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > x {
			x = nums[i]
		} else {
			x += nums[i]
		}
		res = max(res, x)
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

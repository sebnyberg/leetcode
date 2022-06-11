package p2293minmaxgame

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		for i := 0; i < len(nums); i += 2 {
			if i%4 == 0 {
				nums[i/2] = min(nums[i], nums[i+1])
			} else {
				nums[i/2] = max(nums[i], nums[i+1])
			}
		}
		nums = nums[:len(nums)/2]
	}
	return nums[0]
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

package p2760longestevenoddsubarraywiththreshold

func longestAlternatingSubarray(nums []int, threshold int) int {
	queue := []int{}
	var res int
	for i := range nums {
		if nums[i] > threshold {
			queue = queue[:0]
			continue
		}
		if nums[i]&1 == len(queue)&1 {
			queue = append(queue, nums[i])
			res = max(res, len(queue))
			continue
		} else if nums[i]&1 == 0 {
			queue = append(queue[:0], nums[i])
		} else {
			queue = queue[:0]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

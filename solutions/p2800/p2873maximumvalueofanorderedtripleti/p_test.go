package p2873maximumvalueofanorderedtripleti

func maximumTripletValue(nums []int) int64 {
	// Can just exhaustively search
	res := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				res = max(res, (nums[i]-nums[j])*nums[k])
			}
		}
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

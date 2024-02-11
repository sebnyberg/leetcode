package p2932maximumstrongpairxori

func maximumStrongPairXor(nums []int) int {
	var res int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			if abs(a-b) <= min(a, b) {
				x := a ^ b
				if x > res {
					res = x
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

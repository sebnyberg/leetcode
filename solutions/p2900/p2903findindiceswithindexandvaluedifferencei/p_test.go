package p2903findindiceswithindexandvaluedifferencei

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := range nums {
		for j := i; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			if abs(i-j) >= indexDifference && abs(a-b) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

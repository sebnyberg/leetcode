package p3264finalarraystateafterkmultiplicationoperationsi

func getFinalState(nums []int, k int, multiplier int) []int {
	for ; k > 0; k-- {
		var minIdx int
		for i := range nums {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] *= multiplier
	}
	return nums
}

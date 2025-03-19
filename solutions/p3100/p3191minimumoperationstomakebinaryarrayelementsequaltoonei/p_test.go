package p3191minimumoperationstomakebinaryarrayelementsequaltoonei

func minOperations(nums []int) int {
	// You have to flip any one to zero, so just do this in a greedy fashion
	var ops int
	for i := range nums {
		if nums[i] == 1 {
			continue
		}
		if i >= len(nums)-2 {
			return -1
		}
		for k := 0; k < 3; k++ {
			nums[i+k] = 1 - nums[i+k]&1
		}
		ops++
	}
	return ops
}

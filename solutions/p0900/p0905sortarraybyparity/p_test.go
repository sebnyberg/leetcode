package p0905sortarraybyparity

func sortArrayByParity(nums []int) []int {
	// Whenever finding an even number, shuffle it with the current
	// position in nums and increment the position
	var j int
	for i, n := range nums {
		if n%2 == 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return nums
}

package p2367numberofarithmetictriplets

func arithmeticTriplets(nums []int, diff int) int {
	var res int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[k]-nums[j] == diff && nums[j]-nums[i] == diff {
					res++
				}
			}
		}
	}
	return res
}

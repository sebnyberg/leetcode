package p0908smallestrangei

func smallestRangeI(nums []int, k int) int {
	min := nums[0]
	max := nums[0]
	for _, x := range nums {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	d := max - min
	d -= 2 * k
	if d < 0 {
		return 0
	}
	return d
}

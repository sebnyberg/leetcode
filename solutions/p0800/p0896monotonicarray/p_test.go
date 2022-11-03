package p0896monotonicarray

func isMonotonic(nums []int) bool {
	n := len(nums)
	var sign int
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] < nums[i-1] {
			if sign == 1 {
				return false
			}
			sign = -1
			continue
		}
		if sign == -1 {
			return false
		}
		sign = 1
	}
	return true
}

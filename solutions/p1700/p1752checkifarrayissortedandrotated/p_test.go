package p1752checkifarrayissortedandrotated

func check(nums []int) bool {
	var rotated bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if rotated {
				return false
			}
			rotated = true
		}
	}
	return !rotated || nums[len(nums)-1] <= nums[0]
}

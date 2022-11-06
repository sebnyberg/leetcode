package p2460applyoperationstoanarray

func applyOperations(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			res = append(res, nums[i]*2)
			i++
			continue
		}
		res = append(res, nums[i])
	}
	for len(res) < len(nums) {
		res = append(res, 0)
	}
	return res
}

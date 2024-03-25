package p0442findallduplicatesinanarray

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = nums[v-1] * -1
		} else {

			result = append(result, v)
		}
	}
	return result
}

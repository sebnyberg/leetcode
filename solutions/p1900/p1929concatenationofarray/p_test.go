package p1929concatenationofarray

func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, n*2)
	for i := range nums {
		res[i] = nums[i]
		res[i+n] = nums[i]
	}
	return res
}

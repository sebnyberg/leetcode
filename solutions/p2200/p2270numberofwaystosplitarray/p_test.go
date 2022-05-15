package p2270numberofwaystosplitarray

func waysToSplitArray(nums []int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i, num := range nums {
		presum[i+1] = presum[i] + num
	}
	right := nums[n-1]
	var res int
	for i := n - 2; i >= 0; i-- {
		left := presum[i] + nums[i]
		if left >= right {
			res++
		}
		right += nums[i]
	}
	return res
}

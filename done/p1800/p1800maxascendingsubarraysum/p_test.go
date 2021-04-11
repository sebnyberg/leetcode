package p1800maxascendingsubarraysum

func maxAscendingSum(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]
	for i, n := range nums {
		if i == 0 {
			continue
		}
		if n <= nums[i-1] {
			sum = 0
		}
		sum += n
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

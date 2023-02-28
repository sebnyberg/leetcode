package p2574leftandrightsumdifferences

func leftRigthDifference(nums []int) []int {
	n := len(nums)
	left := make([]int, n+1)
	for i := range nums {
		left[i+1] = left[i] + nums[i]
	}
	var right int
	ans := make([]int, n)
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = abs(left[i] - right)
		right += nums[i]
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

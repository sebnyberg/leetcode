package p2614primeindiagonal

func diagonalPrime(nums [][]int) int {
	isprime := func(x int) bool {
		if x <= 1 {
			return false
		}
		for y := 2; y*y <= x; y++ {
			if x%y == 0 {
				return false
			}
		}
		return true
	}

	n := len(nums)
	var res int
	for i := 0; i < n; i++ {
		if isprime(nums[i][i]) {
			res = max(res, nums[i][i])
		}
		if isprime(nums[i][n-i-1]) {
			res = max(res, nums[i][n-i-1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

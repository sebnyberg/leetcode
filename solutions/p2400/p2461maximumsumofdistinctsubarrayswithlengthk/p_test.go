package p2461maximumsumofdistinctsubarrayswithlengthk

func maximumSubarraySum(nums []int, k int) int64 {
	m := make(map[int]int)
	var sum int
	var tooManyCount int
	var maxSum int
	for i, v := range nums {
		sum += v
		if m[v] == 1 {
			tooManyCount++
		}
		m[v]++
		if i >= k {
			sum -= nums[i-k]
			if m[nums[i-k]] == 2 {
				tooManyCount--
			}
			m[nums[i-k]]--
		}
		if i >= k-1 && tooManyCount == 0 {
			maxSum = max(maxSum, sum)
		}
	}
	return int64(maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

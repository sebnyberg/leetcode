package p2913subarraysdistinctelementsumofsquaresi

func sumCounts(nums []int) int {
	// Brute-force
	var res int
	for i := 0; i < len(nums); i++ {
		m := make(map[int]struct{})
		m[nums[i]] = struct{}{}
		res += 1
		for j := i + 1; j < len(nums); j++ {
			m[nums[j]] = struct{}{}
			res += len(m) * len(m)
		}
	}
	return res
}

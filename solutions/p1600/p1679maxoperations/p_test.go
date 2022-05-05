package p1679maxoperations

func maxOperations(nums []int, k int) int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}
	var result int
	for num, count := range counts {
		if num+num == k {
			result += count / 2
		} else {
			result += min(count, counts[k-num])
			counts[num] = 0
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package p2869minimumoperationstocollectelements

func minOperations(nums []int, k int) int {
	seen := make([]bool, k+1)
	var seenCount int
	for i := range nums {
		j := len(nums) - 1 - i
		if nums[j] > k || seen[nums[j]] {
			continue
		}
		seen[nums[j]] = true
		seenCount++
		if seenCount == k {
			return i + 1
		}
	}
	return -1
}

package p3349adjacentincreasingsubarraysdetectioni

func hasIncreasingSubarrays(nums []int, k int) bool {
	var seen bool
	var j int
	for i := range nums {
		if i != j && nums[i] <= nums[j] {
			j = i
		}
		if i-j+1 == k {
			if seen {
				return true
			}
			seen = true
			j = i + 1
		}
	}
	return false
}

package p2918minimumequalsumoftwoarraysafterreplacingzeros

func minSum(nums1 []int, nums2 []int) int64 {
	var sum [2]int
	var zero [2]int

	for _, x := range nums1 {
		sum[0] += x
		if x == 0 {
			zero[0]++
		}
	}
	for _, x := range nums2 {
		sum[1] += x
		if x == 0 {
			zero[1]++
		}
	}

	// Start by filling all zeroes with 1s
	sum[0] += zero[0]
	sum[1] += zero[1]

	if zero[0] == 0 && sum[0] < sum[1] {
		return -1
	}
	if zero[1] == 0 && sum[1] < sum[0] {
		return -1
	}

	return int64(max(sum[0], sum[1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

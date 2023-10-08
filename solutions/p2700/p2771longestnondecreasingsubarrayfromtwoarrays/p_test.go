package p2771longestnondecreasingsubarrayfromtwoarrays

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	var prev [2]int
	var curr [2]int
	res := 1
	prev = [2]int{1, 1}
	for i := 1; i < len(nums1); i++ {
		curr = [2]int{1, 1}
		if nums1[i] >= nums1[i-1] {
			curr[0] = max(curr[0], prev[0]+1)
		}
		if nums1[i] >= nums2[i-1] {
			curr[0] = max(curr[0], prev[1]+1)
		}
		if nums2[i] >= nums1[i-1] {
			curr[1] = max(curr[1], prev[0]+1)
		}
		if nums2[i] >= nums2[i-1] {
			curr[1] = max(curr[1], prev[1]+1)
		}
		res = max(res, curr[0])
		res = max(res, curr[1])

		prev, curr = curr, prev
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

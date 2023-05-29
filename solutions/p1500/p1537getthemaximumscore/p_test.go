package p1537getthemaximumscore

const mod = 1e9 + 7

func maxSum(nums1 []int, nums2 []int) int {
	u := len(nums1) - 1
	l := len(nums2) - 1
	var upper int
	var lower int
	for u >= 0 && l >= 0 {
		for u >= 0 && nums1[u] > nums2[l] {
			upper += nums1[u]
			u--
		}
		if u < 0 {
			break
		}
		for l >= 0 && nums2[l] > nums1[u] {
			lower += nums2[l]
			l--
		}
		if l < 0 {
			break
		}
		if nums1[u] == nums2[l] {
			m := max(upper, lower) + nums1[u]
			upper = m
			lower = m
			u--
			l--
		}
	}
	for u >= 0 {
		upper += nums1[u]
		u--
	}
	for l >= 0 {
		lower += nums2[l]
		l--
	}
	return max(upper, lower) % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

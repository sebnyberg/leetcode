package p2540minimumcommonvalue

func getCommon(nums1 []int, nums2 []int) int {
	var i int
	var j int
	for i < len(nums1) && j < len(nums2) {
		for i < len(nums1) && nums1[i] < nums2[j] {
			i++
		}
		for i < len(nums1) && j < len(nums2) && nums2[j] < nums1[i] {
			j++
		}
		if i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
			return nums1[i]
		}
	}
	return -1
}

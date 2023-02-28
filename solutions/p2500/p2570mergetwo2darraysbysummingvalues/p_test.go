package p2570mergetwo2darraysbysummingvalues

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var i, j int
	var res [][]int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] != nums2[j][0] {
			if nums1[i][0] < nums2[j][0] {
				res = append(res, nums1[i])
				i++
			} else {
				res = append(res, nums2[j])
				j++
			}
		} else {
			res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		}
	}
	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}
	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}
	return res
}

package p1386createtargetarrayinthegivenorder

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i, j := range index {
		res = append(res, 0)
		copy(res[j+1:], res[j:])
		res[j] = nums[i]
	}
	return res
}

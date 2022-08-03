package p2348numberofzerofilledsubarrays

func zeroFilledSubarray(nums []int) int64 {
	var m int
	var res int
	for i := range nums {
		if nums[i] != 0 {
			m = 0
			continue
		}
		m++
		res += m
	}
	return int64(res)
}

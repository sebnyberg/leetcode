package p2527findxorbeautyofarray

func xorBeauty(nums []int) int {
	var res int
	for _, x := range nums {
		res ^= x
	}
	return res
}

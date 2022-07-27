package p2317maximumxorafteroperations

func maximumXOR(nums []int) int {
	var res int
	for _, x := range nums {
		res |= x
	}
	return res
}

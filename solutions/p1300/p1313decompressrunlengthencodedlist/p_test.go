package p1313decompressrunlengthencodedlist

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		for k := 0; k < nums[i]; k++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}

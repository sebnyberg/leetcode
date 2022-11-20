package p2475numberofunequaltripletsinarray

func unequalTriplets(nums []int) int {
	var res int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				a := nums[i]
				b := nums[j]
				c := nums[k]
				if a != b && b != c && a != c {
					res++
				}
			}
		}
	}
	return res
}

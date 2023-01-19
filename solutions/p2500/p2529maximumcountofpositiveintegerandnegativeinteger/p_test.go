package p2529maximumcountofpositiveintegerandnegativeinteger

func maximumCount(nums []int) int {
	var pos, neg int
	for _, x := range nums {
		if x < 0 {
			neg++
		}
		if x > 0 {
			pos++
		}
	}
	if neg < pos {
		return pos
	}
	return neg
}

package p1018binaryprefixdivisibleby5

func prefixesDivBy5(nums []int) []bool {
	var v int
	var res []bool
	for _, x := range nums {
		v = ((v << 1) + x) % 5
		if v == 0 {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

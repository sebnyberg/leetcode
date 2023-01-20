package p4

func maxDepthAfterSplit(seq string) []int {
	// Prefer adding to the lowest-count alternative.
	var first, second int
	var res []int
	for i := range seq {
		if seq[i] == '(' {
			if second < first {
				second++
				res = append(res, 1)
			} else {
				first++
				res = append(res, 0)
			}
		} else {
			if second >= first {
				second--
				res = append(res, 1)
			} else {
				first--
				res = append(res, 0)
			}
		}
	}
	return res
}

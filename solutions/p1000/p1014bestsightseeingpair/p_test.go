package p1014bestsightseeingpair

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	prev := values[0]
	var res int
	for i := 1; i < n; i++ {
		prev--
		res = max(res, values[i]+prev)
		if values[i] > prev {
			prev = values[i]
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
